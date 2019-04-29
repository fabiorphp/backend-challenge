package main

import (
	"github.com/fabiorphp/backend-challenge/pkg/cli"
	ufcli "github.com/urfave/cli"
	"os"
)

var (
	appName = "basket"
	version = "0.0.0"
)

func main() {
	app := ufcli.NewApp()
	app.Name = appName
	app.Version = version
	app.Usage = "Basket Server"
	app.Flags = []ufcli.Flag{
		ufcli.StringFlag{
			Name:   "listen, l",
			Value:  "127.0.0.1:9000",
			Usage:  "Address and port on which Basket Server will accept HTTP requests",
			EnvVar: "LISTEN",
		},
	}

	app.Action = cli.Serve

	_ = app.Run(os.Args)
}
