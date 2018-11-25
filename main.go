package main

import (
	"log"
	"os"

	"github.com/skabbass1/sepoy/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		commands.Schedule,
		commands.Unschedule,
		commands.Info,
		commands.Start,
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
