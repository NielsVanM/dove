package commands

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func TestInitConfigCmdSuccess(t *testing.T) {
	outputDir := "./test.toml"

	err := InitConfigCommand(outputDir, false)
	if err != nil {
		t.Errorf("Init config command failed: %s", err)
		return
	}

	defer os.Remove(outputDir)

	var stat fs.FileInfo
	stat, err = os.Stat(outputDir)
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected config file to be placed in %s but found nil", outputDir)
		return
	}

	expectedSize := len(initConfig)

	if stat.Size() != int64(expectedSize) {
		t.Errorf("Expected content to be of size %d but got %d", expectedSize, stat.Size())
	}
}

func TestInitConfigCmdSuceedOnForce(t *testing.T) {
	outputDir := "./test.toml"

	f, err := os.Create(outputDir)
	if err != nil {
		t.Errorf("Failed to setup test, config not created %s", err)
	}
	f.Close()

	err = InitConfigCommand(outputDir, true)
	defer os.Remove(outputDir)
	if err != nil {
		t.Errorf("Failed to overwrite config, %s", err)
	}
}

func TestInitConfigCmdFailsOnDirectory(t *testing.T) {
	outputDir := "./test.toml"

	err := os.Mkdir(outputDir, fs.FileMode(os.O_RDWR))
	if err != nil {
		t.Errorf("Failed to setup test %s", err)
		return
	}

	err = InitConfigCommand(outputDir, false)

	defer os.Remove(outputDir)
	if !errors.Is(err, ErrConfigIsDirectory) {
		t.Errorf("Unexpected error during config init: %s", err)
	}
}

func TestInitConfigCmdSucceedsOnDirectoryForced(t *testing.T) {
	outputDir := "./test.toml"

	err := os.Mkdir(outputDir, fs.FileMode(os.O_RDWR))
	if err != nil {
		t.Errorf("Failed to setup test %s", err)
		return
	}

	err = InitConfigCommand(outputDir, true)
	if err != nil {
		t.Errorf("Unepxected error: %s", err)
		return
	}

	defer os.Remove(outputDir)
}

func TestInitConfigCmdFailsOnExistingConfig(t *testing.T) {
	outputDir := "./test.toml"

	f, err := os.Create(outputDir)
	if err != nil {
		t.Errorf("Failed to setup test, config not created %s", err)
    return
	}

	f.Close()

	err = InitConfigCommand(outputDir, false)
	defer os.Remove(outputDir)
	if err == nil {
		t.Errorf("No error found when trying to write a config to an already existing project %s", err)
		return
	}

	if !errors.Is(err, ErrConfigExists) {
		t.Errorf("Unexpected error during writing config file %s", err)
	}
}
