package investec

import (
	"context"
	"testing"

	"github.com/nickcorin/investec/mock"
	"github.com/nickcorin/snorlax"
	"github.com/stretchr/testify/suite"
)

type AccessTokenTestSuite struct {
	suite.Suite
	client Client
	server *mock.Server
}

func (suite *AccessTokenTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = NewClient(
		WithTransport(
			snorlax.NewClient(snorlax.WithBaseURL(suite.server.URL)),
		),
	)
}

func (suite *AccessTokenTestSuite) TestClient_GetAccessToken() {
	token, err := suite.client.GetAccessToken(context.TODO(), TokenScopeAccounts)
	suite.Require().NoError(err)
	suite.Require().NotNil(token)

	testToken := AccessToken{
		Token:     "Ms9OsZkyrhBZd5yQJgfEtiDy4t2c",
		Type:      TokenTypeBearer,
		ExpiresIn: 1799,
		Scope:     TokenScopeAccounts,
	}

	suite.Require().Equal(&testToken, token)
}

func TestAccessTokenTestSuite(t *testing.T) {
	suite.Run(t, new(AccessTokenTestSuite))
}
