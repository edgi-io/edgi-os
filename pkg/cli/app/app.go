package app

import (
	"fmt"

	"github.com/edgi-io/edgi-os/pkg/cli/config"
	"github.com/edgi-io/edgi-os/pkg/cli/install"
	"github.com/edgi-io/edgi-os/pkg/cli/rc"
	"github.com/edgi-io/edgi-os/pkg/cli/upgrade"
	"github.com/edgi-io/edgi-os/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	Debug bool
)

// New CLI App
func New() *cli.App {
	app := cli.NewApp()
	app.Name = "edgi"
	app.Usage = "Booting to EDGI so you don't have to"
	app.Version = version.Version
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
	}
	// required flags without defaults will break symlinking to exe with name of sub-command as target
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Turn on debug logs",
			EnvVar:      "EDGI_DEBUG",
			Destination: &Debug,
		},
	}

	app.Commands = []cli.Command{
		rc.Command(),
		config.Command(),
		install.Command(),
		upgrade.Command(),
		//net.Command(),
		//tailscale.Command(),
	}

	app.Before = func(c *cli.Context) error {
		if Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	return app
}
