package ziggy_test

import (
	"context"
	"testing"

	"github.com/nickcorin/ziggy"
	"github.com/nickcorin/ziggy/mock"
	"github.com/stretchr/testify/suite"
)

type AccessTokenTestSuite struct {
	suite.Suite
	client *ziggy.Client
	server *mock.Server
}

func (suite *AccessTokenTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = ziggy.NewClientForTesting(suite.T(), suite.server.URL)
}

func (suite *AccessTokenTestSuite) TestClient_GetAccessToken() {
	token, err := suite.client.GetAccessToken(context.TODO(),
		ziggy.TokenScopeAccounts)
	suite.Require().NoError(err)
	suite.Require().NotNil(token)

	testToken := ziggy.AccessToken{
		Token:     "Ms9OsZkyrhBZd5yQJgfEtiDy4t2c",
		Type:      ziggy.TokenTypeBearer,
		ExpiresIn: 1799,
		Scope:     ziggy.TokenScopeAccounts,
	}

	suite.Require().Equal(testToken.Token, token.Token)
	suite.Require().Equal(testToken.Type, token.Type)
	suite.Require().Equal(testToken.ExpiresIn, token.ExpiresIn)
	suite.Require().Equal(testToken.Scope, token.Scope)

}

func TestAccessTokenTestSuite(t *testing.T) {
	suite.Run(t, new(AccessTokenTestSuite))
}
