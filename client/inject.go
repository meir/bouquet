package client

import (
	"io/fs"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/meir/bouquet/pkg/asar"
	"github.com/meir/bouquet/pkg/discord"
)

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

// applyVersionPatches will apply the patch to inject bouquet scripts properly.
// this will use ./patches/[OS]/[Discord Version]/*.patch
func (c *Client) applyVersionPatches() error {
	os := runtime.GOOS
	version, err := discord.GetVersion()
	if err != nil {
		return err
	}

	return fs.WalkDir(assets, "patches/"+os+"/"+version, func(path string, d fs.DirEntry, err error) error {
		if d == nil || d.IsDir() {
			return nil
		}

		data, err := fs.ReadFile(assets, path)
		if err != nil {
			return err
		}

		ext := filepath.Ext(path)
		if ext == ".patch" {
			return c.applyPatch(data)
		}

		return nil
	})
}

// override_header will overwrite all the files provided in the embedded filesystem or just add them to the asar file
func (c *Client) override_header() error {
	root := c.asarFile.Header
	basePath := "src"

	return fs.WalkDir(assets, basePath, func(path string, d fs.DirEntry, err error) error {
		path = strings.TrimPrefix(path, basePath)
		name := filepath.Base(path)
		ext := filepath.Ext(name)
		parent := filepath.Dir(path)
		folder := root.Get(parent)
		// it should never call this since the root folder is always available.
		// if this error occurs then the dir creation in the next bit somehow failed
		if folder == nil {
			panic("folder is nil: " + parent)
		}

		if d.IsDir() {
			// Create the dir if it does not exist
			if root.Get(path) == nil {
				folder.Add(name, asar.NewFolder())
			}
		} else {
			fullpath := filepath.Join(basePath, "/", path)
			data, err := fs.ReadFile(assets, fullpath)
			if err != nil {
				return err
			}

			// if file is a patch, apply the patch to the ASAR file
			// this might be overwritten if the src/ directory contains a file
			// in the location where the patch is going to be applied
			if ext == ".patch" {
				return c.applyPatch(data)
			}

			// Overwrite the file, if file doesnt exist, create new header for it
			if f := root.Get(path); f != nil {
				f.SetContent(data)
			} else {
				f := asar.NewFile(data, false)
				folder.Add(name, f)
			}
		}
		return nil
	})
}
