package main

import (
	"fmt"
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Dory"
	app.Version = "18.11.1"
	app.Usage = "Ethereum BlockNumber Syncer for kafka."

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Starting Blockchain BlockNumber Syncer Service...\n")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
