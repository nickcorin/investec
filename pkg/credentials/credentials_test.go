package credentials_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/nickcorin/ziggy/pkg/credentials"

	"github.com/stretchr/testify/suite"
)

type credentialsTestSuite struct {
	suite.Suite
}

func TestCredentialsTestSuite(t *testing.T) {
	suite.Run(t, new(credentialsTestSuite))
}

func (suite *credentialsTestSuite) TestStoreGet() {
	id, secret := "ziggy", "s3cr3t"
	err := credentials.Store(id, secret)
	suite.Require().NoError(err)

	creds, err := credentials.Get()
	suite.Require().NoError(err)
	suite.Require().NotNil(creds)
	suite.Require().EqualValues(id, creds.Username)
	suite.Require().EqualValues(secret, creds.Secret)
}

func (suite *credentialsTestSuite) TestPrompt() {
	id, secret := "ziggy", "s3cr3t"

	r, w, err := os.Pipe()
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r

	go func() { fmt.Fprintf(w, "%s\n%s\n", id, secret); w.Close() }()

	creds, err := credentials.Prompt()
	suite.Require().NoError(err)
	suite.Require().NotNil(creds)

	suite.Require().EqualValues(id, creds.Username)
	suite.Require().EqualValues(secret, creds.Secret)
}
