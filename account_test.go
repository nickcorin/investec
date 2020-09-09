package ziggy_test

import (
	"context"
	"testing"

	"github.com/nickcorin/investec/mock"
	"github.com/nickcorin/ziggy"
	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite
	client *ziggy.Client
	server *mock.Server
}

func (suite *AccountTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = ziggy.NewClientForTesting(suite.T(), suite.server.URL)
}

func (suite *AccountTestSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *AccountTestSuite) TestClient_GetAccounts() {
	accounts, err := suite.client.GetAccounts(context.TODO())
	suite.Require().NoError(err)
	suite.Require().NotNil(accounts)

	testAccounts := []ziggy.Account{
		{
			ID:        "172878438321553632224",
			Number:    "10010206147",
			Name:      "Mr John Doe",
			Reference: "My Investec Private Bank Account",
			Product:   "Private Bank Account",
		},
	}

	for _, account := range testAccounts {
		suite.Require().Contains(accounts, account)
	}
}

func (suite *AccountTestSuite) TestClient_GetAccountBalance() {
	balance, err := suite.client.GetAccountBalance(context.TODO(), "123456789")
	suite.Require().NoError(err)
	suite.Require().NotNil(balance)

	testBalance := ziggy.Balance{
		AccountID: "172878438321553632224",
		Current:   28857.76,
		Available: 98857.76,
		Currency:  "ZAR",
	}

	suite.Require().Equal(&testBalance, balance)
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
