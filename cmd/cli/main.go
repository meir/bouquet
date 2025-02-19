package main

import (
	"fmt"
	"log/slog"
	"os"

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
		Version:              bouquet.VERSION,
		EnableBashCompletion: true,

		Commands: []*cli.Command{
			{
				Name:        "install",
				Aliases:     []string{"i"},
				Description: "Backup and install the bouquet in Discord",

				Action: func(c *cli.Context) error {
					asar_path, err := discord.GetPath()
					if err != nil {
						slog.Error("could not find discord location: " + err.Error())
					}

					if err := bouquet.Backup(asar_path); err != nil {
						slog.Error(err.Error())
					}

					if err := bouquet.Restore(asar_path); err != nil {
						slog.Error(err.Error())
					}

					return bouquet.Install(asar_path)

				},
			},
			{
				Name:        "revert",
				Aliases:     []string{"restore", "uninstall", "u"},
				Description: "Revert changes made on Discord by restoring the backup",

				Action: func(c *cli.Context) error {
					asar_path, err := discord.GetPath()
					if err != nil {
						slog.Error("could not find discord location: " + err.Error())
					}

					return bouquet.Restore(asar_path)
				},
			},
			{
				Name:        "extract",
				Aliases:     []string{"e"},
				Description: "Extract the source code from the asar file",

				Action: func(c *cli.Context) error {

					return nil
				},
			},
			{
				Name:        "version",
				Aliases:     []string{"v"},
				Description: "get the installed version of bouquet and the binary version",

				Action: func(c *cli.Context) error {
					fmt.Printf(version_text, bouquet.VERSION)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error(err.Error())
	}
}
