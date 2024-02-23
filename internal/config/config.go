package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type DoveService struct {
  Name string
  Command string
}

type DoveOptions struct {
  Version int

}

type DoveConfig struct {
  DoveOpts DoveOptions `toml:"dove"`
  Services []DoveService `toml:"service"`
}

type ParseConfig func(sources string) (*DoveConfig, error)

func ParseTomlConfig(source string) (*DoveConfig, error) {
  var cfg DoveConfig
  if _, err := toml.DecodeFile(source, &cfg); err != nil {
    fmt.Println(err)
    return  nil, err
  }

  fmt.Println(cfg)


  return &cfg, nil
}

