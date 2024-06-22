package asar

import (
	"io"
	"os"
)

// Asar is a struct that represents an asar file
// An asar file is constructed as followed
// Meta: 16 bytes at the start of the file representing 4 uint32's
// Header: A json object listing all files/folders like a registry
// Content: The actual content of the files, each are defined in the header using size and offset
type Asar struct {
	Location string
	raw      []byte

	Meta   *Meta
	Header *Header
}

// meta size is the size of the meta bytes at the start of an asar file
const META_SIZE = 16

// NewAsar parses an asar file given using the path.
func NewAsar(filePath string) (*Asar, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	meta, err := getMeta(data[:META_SIZE])
	if err != nil {
		return nil, err
	}

	headerEnd := META_SIZE + meta.ContentOffset
	headerData := data[META_SIZE:headerEnd]
	contentData := data[headerEnd:] // suddenly the asar has changed and needs +2 bytes for some reason?

	header, err := getHeader(headerData, contentData)

	return &Asar{
		Location: filePath,
		raw:      data,
		Meta:     meta,
		Header:   header,
	}, nil
}

// Pack will pack the asar file back together at the Location
func (asar *Asar) Pack() (int, error) {
	header, content, err := asar.Header.Pack()
	if err != nil {
		return 0, err
	}

	// update meta using weird magic numbers
	asar.Meta.MetaSize = 4
	asar.Meta.HeaderBufferLength = uint32(len(header)) + 8
	asar.Meta.HeaderLength = uint32(len(header)) + 4
	asar.Meta.ContentOffset = uint32(len(header))

	meta := asar.Meta.ToBytes()

	asarContent := append(meta, append(header, content...)...)

	f, err := os.OpenFile(asar.Location, os.O_RDWR, 0755)
	if err != nil {
		return 0, err
	}

	return f.Write(asarContent)
}
