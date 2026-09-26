package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	tempHomeDir := t.TempDir()
	configTestFile := filepath.Join(tempHomeDir, configFileName)
	t.Setenv("HOME", tempHomeDir)
	err := os.WriteFile(configTestFile, []byte(`
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
	`), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	config, err := Read()
	if err != nil {
		t.Errorf("Could not read file: %v", err)
	}

	if config.DBURL != "connection_string_goes_here" {
		t.Errorf("DBURL should be connection_string_goes_here")
	}

	if config.CurrentUserName != "username_goes_here" {
		t.Errorf("CurrentUserName should be username_goes_here")
	}
}

func TestSetUser(t *testing.T) {
	tempHomeDir := t.TempDir()
	configTestFile := filepath.Join(tempHomeDir, configFileName)
	t.Setenv("HOME", tempHomeDir)
	err := os.WriteFile(configTestFile, []byte(`
{
  "db_url": "connection_string_goes_here"
}
	`), 0o600)
	if err != nil {
		t.Fatal(err)
	}

	config, err := Read()
	if err != nil {
		t.Errorf("Could not read file: %v", err)
	}

	if config.DBURL != "connection_string_goes_here" {
		t.Errorf("DBURL should be connection_string_goes_here")
	}

	if config.CurrentUserName != "" {
		t.Errorf("CurrentUserName should be unset")
	}

	config.SetUser("username_goes_here")

	config, err = Read()
	if err != nil {
		t.Errorf("Could not read file: %v", err)
	}

	if config.DBURL != "connection_string_goes_here" {
		t.Errorf("DBURL should be connection_string_goes_here")
	}

	if config.CurrentUserName != "username_goes_here" {
		t.Errorf("CurrentUserName should be username_goes_here")
	}
}
