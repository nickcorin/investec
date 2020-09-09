package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const credentialsFilename = "ziggy.json"

var unixFilepaths = []string{
	filepath.Join(os.Getenv("$HOME"), "ziggy", credentialsFilename),
	filepath.Join(os.Getenv("XDG_CONFIG_HOME"), credentialsFilename),
	filepath.Join(os.Getenv("HOME"), credentialsFilename),
}

var windowsFilepaths = []string{
	filepath.Join(os.Getenv("systemdrive"), os.Getenv("homepath"),
		credentialsFilename),
}

// Credentials contains the required credentials to authenticate ziggy to access
// your Investec account.
type Credentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

func loadCredentials(paths []string) (*Credentials, error) {
	for _, path := range paths {
		creds, err := loadCredentialsFromFile(path)
		if err != nil {
			continue
		}

		return creds, nil
	}

	return nil, ErrNoCredentials
}

func loadCredentialsFromFile(path string) (*Credentials, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// File doesn't exist.
		return nil, fmt.Errorf("file does not exist %s: %w", path, err)
	} else if err != nil {
		// Unexpected error.
		return nil, fmt.Errorf("failed to open file %s: %w", path, err)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read credentials from %s: %w", path,
			err)
	}

	var credentials Credentials
	if err = json.Unmarshal(data, &credentials); err != nil {
		return nil, fmt.Errorf("failed to unmarshal credentials: %w", err)
	}

	return &credentials, nil
}
