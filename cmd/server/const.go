package main

import (
	"fmt"

	"github.com/eyepipe/eye/internal/pkg/buildinfo"
	"github.com/urfave/cli/v3"
)

var FlagConfigRequired = &cli.StringFlag{
	Name:      "config",
	Value:     "config/config.yml",
	Usage:     fmt.Sprintf("https://github.com/eyepipe/eye/blob/%s/config/config.example.yml", buildinfo.BuildArgGitCommit),
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
