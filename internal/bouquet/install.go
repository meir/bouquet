package bouquet

import (
	"github.com/meir/bouquet/client"
	"github.com/meir/bouquet/pkg/asar"
)

// Install injects the bouquet into the Discord ASAR file after backing it up within the discord app as _core.asar
func Install(asarPath string) error {
	a, err := asar.NewAsar("")
	if err != nil {
		return err
	}

	c := client.NewClient(a)
	if err := c.Build(); err != nil {
		return err
	}

	a.Location = asarPath

	_, err = a.Pack()
	if err != nil {
		panic(err)
	}

	return nil
}
