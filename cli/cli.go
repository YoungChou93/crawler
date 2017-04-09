package cli

import (
	"os"
	"github.com/urfave/cli"
	"path"
	"log"
)

func Run(){

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "A Crawler test"
	app.Version = "0.1.1"

	app.Author = ""
	app.Email = ""


	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
