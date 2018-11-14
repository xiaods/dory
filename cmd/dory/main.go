package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xiaods/dory/pkg/crawler"
	"github.com/xiaods/dory/pkg/version"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Dory"
	app.Version = fmt.Sprintf("%s", version.VERSION)
	app.Usage = "Blockchain Transaction Syncer for kafka."

	app.Action = func(c *cli.Context) error {
		log.Println("Starting Blockchain BlockNumber Syncer Service...")
		// crawler.TrackingETHContract()
		crawler.TrackingETHBalance()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
