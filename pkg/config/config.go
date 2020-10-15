package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// DefaultConfigFile is the default name for Ziggy's config file.
const DefaultConfigFile = ".ziggy.conf"

var defaultConfig = Config{
	PrivacyMask: true,
}

// Config contains the various configuration settings for ziggy.
type Config struct {
	PrivacyMask bool `json:"privacyMask"`
}

func DefaultPath(filename string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to evaluate home directory: %w", err)
	}

	fPath, err := filepath.Abs(filepath.Join(home, filename))
	if err != nil {
		return "", fmt.Errorf("failed to evaluate config path: %w", err)
	}

	return fPath, nil
}

// Clean removes the Ziggy config file if there is one present.
func Clean(filename string) error {
	fPath, err := DefaultPath(filename)
	if err != nil {
		return err
	}

	_, err = os.Stat(fPath)
	if os.IsNotExist(err) {
		// Nothing to do.
		return nil
	} else if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	return os.Remove(fPath)
}

// Load loads and parses the Ziggy config file from disk.
func Load(filename string) (*Config, error) {
	fPath, err := DefaultPath(filename)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	var config Config
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// WriteDefault writes the default config to disk.
func WriteDefault(filename string, overwrite bool) error {
	fPath, err := DefaultPath(filename)
	if err != nil {
		return err
	}

	_, err = os.Stat(fPath)
	if err == nil {
		if overwrite {
			if err = os.Remove(fPath); err != nil {
				return fmt.Errorf("failed to overwrite current config: %w", err)
			}
		} else {
			// I considered returning an error here instead of failing silently,
			// although I think that if a config file alreafy exists and you
			// have chosen not to overwrite - then you should consider the
			// possibility that there may already be a config file present.
			return nil
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to check if config exists: %w", err)
	}

	data, err := json.Marshal(defaultConfig)
	if err != nil {
		return fmt.Errorf("failed to parse default config: %w", err)
	}

	if err = ioutil.WriteFile(fPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write default config: %w", err)
	}

	return nil
}
