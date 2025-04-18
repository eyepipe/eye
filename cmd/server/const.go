package main

import "github.com/urfave/cli/v3"

var FlagConfigRequired = &cli.StringFlag{
	Name:      "config",
	Value:     "config/config.yml",
	TakesFile: true,
}

var FlagAddress = &cli.StringFlag{
	Name:  "address",
	Value: ":3000",
}

var FlagMigrationsDir = &cli.StringFlag{
	Name:  "dir",
	Value: "db/migrations",
}
