package credentials

import "fmt"

var (
	// ErrCredentialsNotFound is returned when retrieving credentials that have
	// not yet been saved.
	ErrCredentialsNotFound = fmt.Errorf("credentials not found")

	// ErrCredentialStoreNotSupported is returned when a suitable credential
	// store for a particular OS is not supported yet.
	ErrCredentialStoreNotSupported = fmt.Errorf("credential store not supported")
)
