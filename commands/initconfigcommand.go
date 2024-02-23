package commands

import (
	_ "embed"
)

//go:embed init.toml
var initConfig string

func InitConfigCommand(targetpath string) error {
  

	return nil
}
