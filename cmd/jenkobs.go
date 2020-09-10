package main

import (
	"os"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/isbm/go-nanoconf"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Setup logger level
func setLogger(ctx *cli.Context) {
	var level logrus.Level
	if ctx.Bool("debug") {
		level = logrus.DebugLevel
	} else {
		level = logrus.InfoLevel
	}
	wzlib_logger.GetCurrentLogger().SetLevel(level)
	wzlib_logger.GetCurrentLogger().Debug("Using debug mode")
}

func listenOBS(ctx *cli.Context) error {
	setLogger(ctx)

	return nil
}

func main() {
	appname := "jenkobs"
	confpath := nanoconf.NewNanoconfFinder(appname).DefaultSetup(nil)
	app := &cli.App{
		Version: "0.1 Alpha",
		Name:    appname,
		Usage:   "AMQP reactor for Open Build Service",
		Action:  listenOBS,
	}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Turn logging debug level",
		},
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Configuration file for actions",
			Value:   confpath.SetDefaultConfig(confpath.FindFirst()).FindDefault(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		wzlib_logger.GetCurrentLogger().Errorf("General error: %s", err.Error())
	}
}
