package asar

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Header contains the entire structure of folders and files.
// If the object is a folder, it will have Files filled with sub headers
// If the object is a file, it will have a size and offset which relate to the
// location and the amount of bytes of the files content within the asar file
type Header struct {
	content []byte `json:"-"`
	name    string

	Size       uint64             `json:"size"`
	Offset     string             `json:"offset,omitempty"`
	Executable bool               `json:"executable,omitempty"`
	Files      map[string]*Header `json:"files,omitempty"`
}

// NewFile creates a new file with the content
func NewFile(content []byte, executable bool) *Header {
	return &Header{
		content: content,

		Size:       uint64(len(content)),
		Offset:     "",
		Executable: executable,
		Files:      map[string]*Header{},
	}
}

// NewFolder creates a new empty folder
func NewFolder() *Header {
	return &Header{
		Files: map[string]*Header{},
	}
}

// Add adds a new file or folder under a folder
// if used on a file it will return an error
func (f *Header) Add(name string, h *Header) error {
	if f.Size != 0 {
		return fmt.Errorf("cannot add files/folders under files")
	}

	h.name = name
	f.Files[name] = h
	return nil
}

// getHeader will parse the header json object and load the file content into the header as Content
func getHeader(register []byte, content []byte) (*Header, error) {
	var fileRegister *Header = &Header{}
	err := json.Unmarshal(register, fileRegister)
	if err != nil {
		panic(err)
	}

	fileRegister.load("root", content)

	return fileRegister, nil
}

// load will load the content into the header object recursively for all the files/folders in the header
func (f *Header) load(name string, content []byte) {
	f.name = name
	if len(f.Files) > 0 {
		for name, file := range f.Files {
			file.load(name, content)
		}
	} else if f.Size > 0 {
		offset, err := strconv.Atoi(f.Offset)
		if err != nil {
			panic("offset is not a number")
		}
		f.content = content[offset : uint64(offset)+f.Size]
	}
}

// Name returns the name of the file/folder
func (f *Header) Name() string {
	return f.name
}

// Content returns the file content as a byte array
func (f *Header) Content() []byte {
	return f.content
}

// SetContent will set the content of the file
func (f *Header) SetContent(content []byte) {
	f.content = content
	f.Size = uint64(len(content))
}

// Get will return the header object for the given path
func (f *Header) Get(path string) *Header {
	path = strings.TrimPrefix(path, "/")
	if path == "" || path == "." {
		return f
	}
	p := strings.Split(path, "/")
	return f.get(p)
}

// get will return the header object for the given path
func (f *Header) get(path []string) *Header {
	if len(path) == 0 {
		return f
	}

	p := path[0]
	if f.Files == nil {
		return nil
	}

	if sf, ok := f.Files[p]; ok && sf != nil {
		return sf.get(path[1:])
	}
	return nil
}

// Pack will pack the header and content into a byte array and return these
func (f *Header) Pack() ([]byte, []byte, error) {
	content := f.packContent([]byte{})
	header, err := json.Marshal(f)
	return header, content, err
}

// packContent will pack all the content recursively into the byte array and update the header with the new offset and size
func (f *Header) packContent(content []byte) []byte {
	if len(f.Files) > 0 {
		for _, file := range f.Files {
			content = file.packContent(content)
		}
		return content
	}

	f.Offset = strconv.Itoa(len(content))
	f.Size = uint64(len(f.content))

	content = append(content, f.content...)
	return content
}
