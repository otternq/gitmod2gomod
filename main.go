// Command gitmod2gomod converts git submodules in a vendor directory to a go.mod file.
package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var (
		err error
		app = cli.NewApp()
	)

	app.Name = "gitmod2gomod"
	app.Usage = "Convert git submodule vendor dependencies to go mod format"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "repo-path",
			Usage: "The path to the git repository",
		},
	}

	if err = app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
