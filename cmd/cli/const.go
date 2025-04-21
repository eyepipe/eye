package main

import (
	"time"

	"github.com/eyepipe/eye/internal/pkg/container"
	"github.com/urfave/cli/v3"
)

var DefaultSchemeJSON = MustMarshalJSON(container.SchemeDefault)

var IFlag = &cli.StringFlag{
	Name:      "i",
	Aliases:   []string{"identity"},
	Usage:     "your private key file (keep it in secret)",
	TakesFile: true,
}

var PFlag = &cli.StringFlag{
	Name:      "p",
	Aliases:   []string{"participant"},
	Usage:     "the public key of your counterpart/recipient",
	TakesFile: true,
}
var ContractURLFlag = &cli.StringFlag{
	Name:    "contract-url",
	Value:   "https://api.eyepipe.pw/v1",
	Sources: cli.EnvVars("EYE_CONTRACT_URL"),
}

var FlagScheme = &cli.StringFlag{
	Name:    "s",
	Aliases: []string{"scheme"},
	Value:   string(DefaultSchemeJSON),
	Usage: `a URL or JSON containing the schema of the encryption 
and signature algorithms that will be used with this key.

Available defaults:

- https://api.eyepipe.pw/v1/schemes/super.json
- https://api.eyepipe.pw/v1/schemes/high.json

`,
}

var FlagSigHex = &cli.StringFlag{
	Name: "sig-hex",
}

var FlagSig = &cli.StringFlag{
	Name:      "sig",
	TakesFile: true,
}

var FlagTimeout7s = &cli.DurationFlag{
	Name:  "timeout",
	Usage: "entire operation and/or HTTP request timeout",
	Value: 7 * time.Second,
}
