package commands

import (
	_ "embed"
	"errors"
	"os"
	"strings"
)

var ErrConfigExists = errors.New("Config file already exists")
var ErrConfigIsDirectory = errors.New("Config file is a directory")

//go:embed init.toml
var initConfig string

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

	f.Write([]byte(initConfig))

	return nil
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
