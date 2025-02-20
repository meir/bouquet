package client

import (
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/meir/bouquet/pkg/asar"
)

// client_root is the root directory of the client files in the ASAR file and in src
const client_root = "app_bootstrap"

// Client is the client injector for bouquet
type Client struct {
	asarFile    *asar.Asar
	buildResult api.BuildResult
}

// NewClient creates a new client injector
func NewClient(asarFile *asar.Asar) *Client {
	return &Client{
		asarFile: asarFile,
	}
}

// client_header will return the asar header for the built client
// will return nil if build has not been ran or has no output files
func (c *Client) client_header() *asar.Header {
	if c.buildResult.OutputFiles == nil {
		return nil
	}

	header := asar.NewFolder()

	for _, file := range c.buildResult.OutputFiles {
		path := filepath.Base(file.Path)
		fileHeader := asar.NewFile(file.Contents, false)
		header.Add(path, fileHeader)
	}

	return header
}

// Build will write the bouquet client into the provided asar file
// to run the asar file, you will still need to pack it and place it at the correct location
func (c *Client) Build() error {
	// build the client from src/bouquet
	if err := c.build(); err != nil {
		return err
	}

	// creates a header out of the built client
	client := c.client_header()

	// add the client to the asar file
	c.asarFile.Header.Add(client_root, client)

	return nil
}
