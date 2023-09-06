package EasyGoQ_Example

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "EasyGoQ Example",
		Usage: "publisher",
		Commands: []*cli.Command{
			{
				Name:    "publisher",
				Aliases: []string{"publish"},
				Usage:   "",
				Action: func(c *cli.Context) error {

					return nil
				},
			},
			{
				Name:    "consumer",
				Aliases: []string{"consume"},
				Usage:   "",
				Action: func(c *cli.Context) error {

					return nil
				},
			},
		},
		Action: func(*cli.Context) error {
			//publish

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
