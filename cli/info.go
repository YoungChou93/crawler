package cli

import (
	"github.com/urfave/cli"
	"log"
)

func info(c *cli.Context){
	if len(c.Args()) != 0 {
		log.Fatalf("the `create` command takes no arguments. See '%s create --help'.", c.App.Name)
	}

	aCrawler.Print()
}