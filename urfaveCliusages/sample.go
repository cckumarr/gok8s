package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// you cn see
func main() {
	var language string
	var dragon string
	app := cli.NewApp()
	var printString string
	var greeting string
	
	//the name value etc on the left is what the cli requires,
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for the greeting",
			Destination: &language,
		},
		cli.StringFlag{
			Name:        "dragon",
			Value:       "english",
			Usage:       "are you a dragon ?",
			Destination: &dragon,
		},
	}

	app.Name = "sayhello"
	app.Usage = "this says hello"

	app.Action = func(c *cli.Context) error {
		name := "there"
		if c.NArg() > 0 {
			name = c.Args()[0]
		}

		if language == "spanish" {
			greeting = "hola"
		} else {
			greeting = "hello"
		}
		
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
