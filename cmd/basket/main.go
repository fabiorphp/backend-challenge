package main

import (
	"fmt"
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
	app.Usage = "Basket App"

	hostFlag := ufcli.StringFlag{
		Name:  "host",
		Value: "http://127.0.0.1:9000",
		Usage: "Address and port of Basket Server",
	}

	app.Commands = []ufcli.Command{
		ufcli.Command{
			Name:   "agent",
			Usage:  "Agent server",
			Action: cli.Serve,
			Flags: []ufcli.Flag{
				ufcli.StringFlag{
					Name:   "listen, l",
					Value:  "127.0.0.1:9000",
					Usage:  "Address and port on which Agent server will accept HTTP requests",
					EnvVar: "LISTEN",
				},
			},
		},
		ufcli.Command{
			Name:   "create",
			Usage:  "Create basket",
			Action: cli.Create,
			Flags:  []ufcli.Flag{hostFlag},
		},
		ufcli.Command{
			Name:   "delete",
			Usage:  "Delete basket",
			Action: cli.Delete,
			Flags:  []ufcli.Flag{hostFlag},
		},
		ufcli.Command{
			Name:   "add",
			Usage:  "Add item into basket",
			Action: cli.Add,
			Flags:  []ufcli.Flag{hostFlag},
		},
		ufcli.Command{
			Name:   "amount",
			Usage:  "Basket amount",
			Action: cli.Amount,
			Flags:  []ufcli.Flag{hostFlag},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%v\n\n", err)
		os.Exit(1)
	}
}
