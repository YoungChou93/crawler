package cli

import "github.com/urfave/cli"

var (
	commands = []cli.Command{
		{
			Name:      "start",
			ShortName: "s",
			Usage:     "Start a task",
			Action:    start,
		},
		{
			Name:      "info",
			ShortName: "i",
			Usage:     "information of crawler",
			Action:    info,
		},
	}

)