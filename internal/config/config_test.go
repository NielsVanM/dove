package config

import "testing"

var expectedNames = []string{
	0: "tailwindcss",
	1: "server",
	2: "frontend",
}

func TestParseTomlConfigParsesOkay(t *testing.T) {
	cfg, err := ParseTomlConfig("../../testdata/config.toml")

	if err != nil {
		t.Errorf("Failed to parse valid toml file: %s", err)
		return
	}

	if cfg.DoveOpts.Version != 1 {
		t.Errorf("Expected version 1 got %d", cfg.DoveOpts.Version)
	}

	if len(cfg.Services) != 3 {
		t.Errorf("Expected 3 services got %d", len(cfg.Services))
	}

	for k, v := range expectedNames {
		if cfg.Services[k].Name != v {
			t.Errorf("Expected Service index %d's name to be %s, got %s", k, v, cfg.Services[k].Name)
		}
	}
}

