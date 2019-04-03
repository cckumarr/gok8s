package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// example:  ./sample --lang "spanish" --dragon "bow to me" mx
// the arguments are captured after the named arguments that is --
// if the name of the string flag is the same as another then it will be a runtime error as soon as the application starts
func main() {
	var language string
	var dragon string
	app := cli.NewApp()
	//var printString string
	var greeting string
	var name string

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
			Value:       "",
			Usage:       "are you a dragon ?",
			Destination: &dragon,
		},
	}

	app.Name = "sayhello"
	app.Usage = "this says hello"

	app.Action = func(c *cli.Context) error {
		
		if c.NArg() > 0 {
			name = c.Args()[0]
		}

		if language == "spanish" {
			greeting = "hola"
		} else {
			greeting = "hello"
		}

		fmt.Println(greeting+ " " + name + " " + dragon)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
