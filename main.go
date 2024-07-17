package main

import(
	"github.com/urfave/cli/v2"
	"fmt"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "Website health checker",
		Usage: "A tiny CLI tool to ping websites",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aleases: []string{"d"},
				Usage: "The domain name to check.",
				Requited: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aleases []string{"p"},
				Usage: "The port to check. Default: 80",
				Requited: false,
			},
		},
		Action: func(c *cli.Context) error{
			port := c.Strings("port")
			if c.Strings("port") == "" {
				port = "80"
			}
			status := Check(c.Strings("domain"), port)
			fmt.PrintLn(status)
			return nil
		},
	
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}