package credentials

import (
	"fmt"
	"os"
	"runtime"

	"github.com/docker/docker-credential-helpers/client"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/nickcorin/ziggy"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	credentialStoreDarwin  = "osxkeychain"
	credentialStoreLinux   = "secretservice"
	credentialStoreWindows = "wincred"
)

func nativeStore() (client.ProgramFunc, error) {
	switch runtime.GOOS {
	case "darwin":
		return client.NewShellProgramFunc(credentialStoreDarwin), nil
	case "linux":
		return client.NewShellProgramFunc(credentialStoreLinux), nil
	case "windows":
		return client.NewShellProgramFunc(credentialStoreWindows), nil
	}

	return nil, ErrCredentialStoreNotSupported
}

// Erase removes any credentials from the OS specific secrets manager.
func Erase() error {
	store, err := nativeStore()
	if err != nil {
		return err
	}

	return client.Erase(store, ziggy.DefaultURL)
}

// Get retrieves credentials from the OS specific secrets manager.
func Get() (*credentials.Credentials, error) {
	store, err := nativeStore()
	if err != nil {
		return nil, err
	}

	creds, err := client.Get(store, ziggy.DefaultURL)
	if credentials.IsErrCredentialsNotFound(err) {
		creds, err = Prompt()
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	err = Store(creds.Username, creds.Secret)
	if err != nil {
		return nil, err
	}

	return creds, nil
}

// Prompt provides a way for users to safely input credentials on the command
// line.
func Prompt() (*credentials.Credentials, error) {
	fmt.Fprint(os.Stdin, "Client ID: ")
	clientBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	fmt.Fprint(os.Stdin, "Client Secret :")
	secretBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	return &credentials.Credentials{
		ServerURL: ziggy.DefaultURL,
		Username:  string(clientBytes),
		Secret:    string(secretBytes),
	}, nil
}

// Store saves the provided credentials in the OS specific secrets manager.
func Store(clientID, clientSecret string) error {
	store, err := nativeStore()
	if err != nil {
		return err
	}

	creds := credentials.Credentials{
		ServerURL: ziggy.DefaultURL,
		Username:  clientID,
		Secret:    clientSecret,
	}

	return client.Store(store, &creds)
}
