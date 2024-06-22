package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	version "github.com/meir/bouquet"
	"github.com/meir/bouquet/internal/bouquet"
	"github.com/meir/bouquet/pkg/discord"

	"github.com/urfave/cli/v2"
)

const version_text = `path: %s
discord host version: %s
installed bouquet version: %s
bouquet binary version: %s
`

func main() {
	app := &cli.App{
		Name:                 "bouquet",
		Version:              version.VERSION,
		EnableBashCompletion: true,

		Commands: []*cli.Command{
			{
				Name:        "install",
				Aliases:     []string{"i"},
				Description: "Backup and install the bouquet in Discord",

				Action: func(c *cli.Context) error {
					path, err := discord.GetPath()
					if err != nil {
						slog.Error("could not find discord location: " + err.Error())
					}
					asarPath := discord.GetASARPath(path)

					if err := bouquet.Backup(asarPath); err != nil {
						slog.Error(err.Error())
					}

					if err := bouquet.Restore(asarPath); err != nil {
						slog.Error(err.Error())
					}

					return bouquet.Install(asarPath)

				},
			},
			{
				Name:        "revert",
				Aliases:     []string{"restore", "uninstall", "u"},
				Description: "Revert changes made on Discord by restoring the backup",

				Action: func(c *cli.Context) error {
					path, err := discord.GetPath()
					if err != nil {
						slog.Error("could not find discord location: " + err.Error())
					}
					asarPath := discord.GetASARPath(path)

					return bouquet.Restore(asarPath)
				},
			},
			{
				Name:        "version",
				Aliases:     []string{"v"},
				Description: "get the installed version of bouquet and the binary version",

				Action: func(c *cli.Context) error {
					path, err := discord.GetPath()
					if err != nil {
						panic("could not find discord location: " + err.Error())
					}
					asarPath := discord.GetASARPath(path)
					discordVersion := filepath.Base(path)

					installed, binary, err := bouquet.Version(asarPath)
					if err != nil {
						return err
					}

					if installed == "" {
						installed = "not installed"
					}

					fmt.Printf(version_text, path, discordVersion, installed, binary)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
	}
}
