package asar

import (
	"bytes"
	"encoding/binary"
)

// Meta contains the metadata of an asar file
type Meta struct {
	MetaSize           uint32
	HeaderBufferLength uint32
	HeaderLength       uint32
	ContentOffset      uint32
}

// getMeta will parse the meta bytes at the start of an asar file
func getMeta(data []byte) (*Meta, error) {
	meta := &Meta{}
	reader := bytes.NewReader(data)

	// Read the version
	err := binary.Read(reader, binary.LittleEndian, &meta.MetaSize)
	if err != nil {
		return nil, err
	}

	// Read the header length
	err = binary.Read(reader, binary.LittleEndian, &meta.HeaderBufferLength)
	if err != nil {
		return nil, err
	}

	// Read the content length
	err = binary.Read(reader, binary.LittleEndian, &meta.HeaderLength)
	if err != nil {
		return nil, err
	}

	// Read the content offset
	err = binary.Read(reader, binary.LittleEndian, &meta.ContentOffset)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

// ToBytes will convert the meta struct into a byte array
func (m *Meta) ToBytes() []byte {
	meta := make([]byte, META_SIZE)
	binary.LittleEndian.PutUint32(meta[0:4], m.MetaSize)
	binary.LittleEndian.PutUint32(meta[4:8], m.HeaderBufferLength)
	binary.LittleEndian.PutUint32(meta[8:12], m.HeaderLength)
	binary.LittleEndian.PutUint32(meta[12:16], m.ContentOffset)
	return meta
}
