package commands

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	DOVE_CFG_FILE_NAME = ".dovecfg"
)

var ErrConfigExists = errors.New("Config file already exists")
var ErrConfigIsDirectory = errors.New("Config file is a directory")

//go:embed init.toml
var initConfig []byte

func InitConfigCommand(targetpath string, force bool) error {
	if force {
		os.Remove(targetpath)
	} else if fileExists(targetpath) {
		return ErrConfigExists
	}

	f, err := os.Create(targetpath)
	if err != nil {
		if strings.Contains(err.Error(), "is a directory") {
			return ErrConfigIsDirectory
		}

		if strings.Contains(err.Error(), "already exists") {
			return ErrConfigExists
		}

		return err
	}

	f.Write(initConfig)

	fmt.Println("Written default config to " + f.Name())

	return nil
}

func ProcessInitArguments(ctx *cli.Context) (string, bool) {
	targetPath := ctx.Args().Get(0)
	if len(targetPath) == 0 {
		targetPath = "./" + DOVE_CFG_FILE_NAME
	}

	if !strings.HasSuffix(targetPath, DOVE_CFG_FILE_NAME) {
		if !strings.HasSuffix(targetPath, "/") {
			targetPath += "/"
		}
		targetPath += DOVE_CFG_FILE_NAME
	}

	force := ctx.Bool("force")

	return targetPath, force
}

func fileExists(filepath string) bool {
	stat, err := os.Stat(filepath)
	if err != nil {
		return false
	}

	if stat == nil {
		return false
	}

	if stat.IsDir() {
		return false
	}

	return true

}
