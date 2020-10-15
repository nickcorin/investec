package credentials

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/nickcorin/ziggy"

	"github.com/docker/docker-credential-helpers/client"
	"github.com/docker/docker-credential-helpers/credentials"
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
	id, err := prompt("Client ID: ")
	if err != nil {
		return nil, err
	}

	secret, err := prompt("Client Secret: ")
	if err != nil {
		return nil, err
	}

	return &credentials.Credentials{
		ServerURL: ziggy.DefaultURL,
		Username:  string(id),
		Secret:    string(secret),
	}, nil
}

func prompt(prompt string) ([]byte, error) {
	fd := int(os.Stdin.Fd())
	if terminal.IsTerminal(fd) {
		fmt.Fprint(os.Stdout, prompt)
		pw, err := terminal.ReadPassword(fd)
		if err != nil {
			return nil, err
		}
		fmt.Fprintln(os.Stdout)
		return pw, nil
	}

	var b [1]byte
	var pw []byte
	for {
		n, err := os.Stdin.Read(b[:])
		// terminal.ReadPassword discards any '\r', so we do the same
		if n > 0 && b[0] != '\r' {
			if b[0] == '\n' {
				return pw, nil
			}
			pw = append(pw, b[0])
			// limit size, so that a wrong input won't fill up the memory
			if len(pw) > 1024 {
				return nil, fmt.Errorf("password too long")
			}
		}
		if err != nil {
			// terminal.ReadPassword accepts EOF-terminated passwords
			// if non-empty, so we do the same
			if err == io.EOF && len(pw) > 0 {
				err = nil
			}
			return pw, err
		}
	}
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
