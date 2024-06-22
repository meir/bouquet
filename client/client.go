package client

import (
	"github.com/evanw/esbuild/pkg/api"
	version "github.com/meir/bouquet"
	"github.com/meir/bouquet/pkg/asar"
)

// client_root is the root directory of the client files in the ASAR file and in src
const client_root = "bouquet"

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

// Inject will inject the bouquet client into the provided asar file
// to run the asar file, you will still need to pack it and place it at the correct location
func (c *Client) Inject() error {
	// apply the patches for the injection hooks
	if err := c.applyVersionPatches(); err != nil {
		return err
	}

	// override headers are the files that just need to be overwritten and not built
	if err := c.override_header(); err != nil {
		return err
	}

	// build the client from src/bouquet
	if err := c.build(); err != nil {
		return err
	}

	// creates a header out of the built client
	client := c.client_header()
	versionFile := asar.NewFile([]byte(version.VERSION), false)
	client.Add("VERSION", versionFile)

	// add the client to the asar file
	c.asarFile.Header.Add(client_root, client)

	return nil
}
