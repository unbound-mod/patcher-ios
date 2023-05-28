package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "patcher-ios",
		Usage: "Patches the Discord IPA with icons, utilities and features to ease usability.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "ipa",
				Aliases: []string{"i"},
				Usage:   "The `path` of the IPA file you would like to patch.",
			},
		},
		Action: func(context *cli.Context) error {
			fmt.Printf("Hello %q", context.Args().Get(0))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
