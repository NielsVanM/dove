package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/nielsvanm/dove/commands"
	"github.com/urfave/cli/v2"
)

// TODO:
// - Config parsing
// - Command starter
// - Process monitor
// - Command onboarding
// - Multi screen with hotkeys to switch between outputs
// - Maybe interactive processes? (opt)
// - Interval runner (opt)

const toolArt = `

       .-''-.
      / ,    \
    .-'(o)    ;
    -==.       |
        \._...-;-.
        )--"""   \-.
        /   .        \-.
      /   /       .    \-.
      |   \    ;   \      \-._________
      |    \    \. .;          -------\.
        \    \-.   \\\\          \---...|
        \.     '-. \\\\.--'._   \---...|
          \-.....7 -.))\     '-._'-.. /
            \._\ /   '-'         '-.,/
              / /
              /=(_
          -./--' 
          ,^-(_
        ,--'
`


const (
  DOVE_CFG_FILE_NAME = ".dovecfg"
  CLI_ERR_INVALID_ARGUMENT = 1
)

func main() {
	app := cli.App{
		Name:           "dove",
		Usage:          "Run multiple processes at the same time for development!" + toolArt,
		DefaultCommand: "start",
		Authors: []*cli.Author{
			{
				Name:  "Niels van Marion",
				Email: "niels.van.marion@quicknet.nl",
			},
		},
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Usage:       "The config to use when starting dev processes",
				Aliases:     []string{"c"},
				DefaultText: "./",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "init",
				Description: "Initializes the local directory with a dove config",
				Args:        true, 
        Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "force",
						Usage:   "Forces the init process even though a dove config is already pressent",
						Aliases: []string{"f"},
					},
				},
				Action: func(ctx *cli.Context) error {
					targetPath := ctx.Args().Get(0)
          if len(targetPath) == 0 {
            targetPath = "./" + DOVE_CFG_FILE_NAME
          }


          if !strings.HasSuffix(targetPath, DOVE_CFG_FILE_NAME) {
            if !strings.HasSuffix(targetPath, "/")  {
              targetPath += "/"
            }
            targetPath += DOVE_CFG_FILE_NAME
          }

          force := ctx.Bool("force")
          fmt.Println(targetPath)

          return commands.InitConfigCommand(targetPath, force)
				},
			},
			{
				Name:        "start",
				Description: "Starts the dev services defined in the dovecfg.toml",
				Aliases:     []string{"dev", "d", "s"},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		slog.Error("Application exited with error", "err", err)
	}
}
