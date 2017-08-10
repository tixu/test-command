package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "19.99.0"

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start a task",
			Action: func(c *cli.Context) error {
				if !c.Args().Present() {
					//return errors.New("no arhime")
					os.Exit(2)
				}
				fmt.Println("started task ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "terminate",
			Aliases: []string{"h"},
			Usage:   "terminate a task",
			Action: func(c *cli.Context) error {
				if !c.Args().Present() {
					//return errors.New("no arhime")
					os.Exit(2)
				}
				fmt.Println("terminate task ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"a"},
			Usage:   "list all the tasks",
			Action: func(c *cli.Context) error {
				fmt.Println("list all the task ")
				return nil
			},
		},
	}
	var hostname string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "hostname",
			Value:       "http://0.0.0.0:12",
			Usage:       "Remote Hostname",
			Destination: &hostname,
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}
	app.Run(os.Args)
}
