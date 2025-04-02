package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "MyApp",
		Usage: "A simple CLI to say hello",
		Commands: []*cli.Command{
			{
				Name:  "hello",
				Usage: "Prints a greeting message",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Value: "Guest", Usage: "Your name"},
				},
				Action: func(c *cli.Context) error {
					name := c.String("name")
					fmt.Printf("Hello, %s!\n", name)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
