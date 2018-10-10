package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kubehippie/database-operator/pkg/action"
	"github.com/kubehippie/database-operator/pkg/config"
	"github.com/kubehippie/database-operator/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	cfg := config.Load()

	if env := os.Getenv("DATABASE_OPERATOR_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:    "database-operator",
		Version: version.Version,
		Usage:   "Database Operator for Kubernetes",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Value:       "info",
				Usage:       "Only log messages with given severity",
				EnvVars:     []string{"DATABASE_OPERATOR_LOG_LEVEL"},
				Destination: &cfg.Logs.Level,
			},
			&cli.BoolFlag{
				Name:        "log-pretty",
				Value:       false,
				Usage:       "Enable pretty messages for logging",
				EnvVars:     []string{"DATABASE_OPERATOR_LOG_PRETTY"},
				Destination: &cfg.Logs.Pretty,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start integrated server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "kubeconfig",
						Value:       "",
						Usage:       "Path to a kubeconfig file",
						EnvVars:     []string{"DATABASE_OPERATOR_SERVER_KUBECONFIG"},
						Destination: &cfg.Server.Kubeconfig,
					},
					&cli.StringFlag{
						Name:        "address",
						Value:       "0.0.0.0:9090",
						Usage:       "Address to bind the server",
						EnvVars:     []string{"DATABASE_OPERATOR_SERVER_ADDRESS"},
						Destination: &cfg.Server.Addr,
					},
				},
				Action: func(c *cli.Context) error {
					return action.Server(cfg, setupLogger(cfg))
				},
			},
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
