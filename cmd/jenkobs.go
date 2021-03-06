package main

import (
	"os"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	"github.com/isbm/go-nanoconf"
	jenkobs_reactor "github.com/isbm/jenkobs/reactor"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var VERSION string = "0.1"

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

	wzlib_logger.GetCurrentLogger().Debug("Connecting to the AMQP")
	conf := nanoconf.NewConfig(ctx.String("config"))

	return jenkobs_reactor.NewReactor().
		SetAMQPAuth(jenkobs_reactor.NewAMQPAuth(conf.Find("amqp"))).
		SetJenkinsAuth(jenkobs_reactor.NewJenkinsAuth(conf.Find("jenkins"))).
		LoadActions(conf.Root().String("actions", "")).
		Run()
}

func main() {
	appname := "jenkobs"
	confpath := nanoconf.NewNanoconfFinder(appname).DefaultSetup(nil)
	app := &cli.App{
		Version: VERSION,
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
