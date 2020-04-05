package main

import (
	"github.com/urfave/cli"
	"goss/Commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "goss"
	app.Usage = "set up oss contributing"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) error {
		Commands.Create(c.Args().Get(0), c.Args().Get(1))
		return nil
	}

	app.Run(os.Args)
}
