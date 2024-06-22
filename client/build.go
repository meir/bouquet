package client

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/evanw/esbuild/pkg/api"
)

// build builds the bouquet client using esbuild
func (c *Client) build() error {
	c.buildResult = api.Build(api.BuildOptions{
		EntryPoints: []string{"src/bouquet/index.ts"},
		Bundle:      true,
		Write:       false,
		Outdir:      "/bouquet",
		Platform:    api.PlatformNode,
		JSXFactory:  "JSX.element",
		JSXFragment: "JSX.fragment",
		Target:      api.ES2020,
		Plugins:     []api.Plugin{c.plugin()},
	})

	if len(c.buildResult.Errors) > 0 {
		// wrap all errors
		var err error = nil
		for _, msg := range c.buildResult.Errors {
			err = errors.Join(err, fmt.Errorf("%s (%v)", msg.Text, msg.Location))
		}
		return err
	}

	return nil
}

// plugin returns an esbuild plugin in order to load the source files from the embedded filesystem
func (c *Client) plugin() api.Plugin {
	return api.Plugin{
		Name: "bouquet",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: ".*"}, c.resolve)
			build.OnLoad(api.OnLoadOptions{Filter: ".*"}, c.load)
		},
	}
}

// resolve will find the correct path for the requested file in the embedded file system
// if file cannot be resolved in the embedded fs it will mark it as an external file
// if the file cannot be found then this will only give an error on runtime
func (c *Client) resolve(args api.OnResolveArgs) (api.OnResolveResult, error) {
	path := filepath.Join(filepath.Dir(args.Importer), args.Path)
	if path := resolvePath(path, false); path != "" {
		return api.OnResolveResult{Path: path, Namespace: "bouquet"}, nil
	}
	return api.OnResolveResult{External: true}, nil
}

// load will load in the file from the embedded filesystem
func (c *Client) load(args api.OnLoadArgs) (api.OnLoadResult, error) {
	data, err := assets.ReadFile(args.Path)
	if err != nil {
		return api.OnLoadResult{}, err
	}

	content := string(data)
	return api.OnLoadResult{
		Contents: &content,
		Loader:   api.LoaderTSX,
	}, nil
}
