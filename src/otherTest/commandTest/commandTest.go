package main

import (
	"Go-algorithm/src/otherTest/commandTest/inner"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("study")

func main() {
	local := []*cli.Command{
        inner.SubCommand,
	}
	app := &cli.App{
		Name:     "local",
		Usage:    "local test",
		Description: "to say somethings",
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Value: "info",
			},
		},
		Before: func(cctx *cli.Context) error {
			return logging.SetLogLevel("study", cctx.String("log-level"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		os.Exit(1)
		return
	}
}