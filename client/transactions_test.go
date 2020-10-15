package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/nickcorin/ziggy"
	"github.com/nickcorin/ziggy/client"
	"github.com/nickcorin/ziggy/server"

	"github.com/stretchr/testify/suite"
)

type TransactionsTestSuite struct {
	suite.Suite
	client ziggy.Client
	server *server.Mock
}

func (suite *TransactionsTestSuite) SetupSuite() {
	suite.server = server.NewMock()
	suite.client = client.NewHTTPForTesting(suite.T(), suite.server.URL)
}

func (suite *TransactionsTestSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *TransactionsTestSuite) TestClient_GetTransactions() {
	res, err := suite.client.GetTransactions(context.TODO(),
		&ziggy.TransactionsRequest{AccountID: "123456789"})
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	transactions := []ziggy.Transaction{
		{
			AccountID:       "172878438321553632224",
			Type:            ziggy.TransactionTypeDebit,
			Status:          ziggy.TransactionStatusPosted,
			Description:     "MONTHLY SERVICE CHARGE",
			CardNumber:      "",
			PostingDate:     time.Date(2020, 6, 11, 0, 0, 0, 0, time.UTC),
			ValueDate:       time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			ActionDate:      time.Date(2020, 6, 18, 0, 0, 0, 0, time.UTC),
			TransactionDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			Amount:          535,
		},
		{
			AccountID:       "172878438321553632224",
			Type:            ziggy.TransactionTypeCredit,
			Status:          ziggy.TransactionStatusPosted,
			Description:     "CREDIT INTEREST",
			CardNumber:      "",
			PostingDate:     time.Date(2020, 6, 11, 0, 0, 0, 0, time.UTC),
			ValueDate:       time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			ActionDate:      time.Date(2020, 6, 18, 0, 0, 0, 0, time.UTC),
			TransactionDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			Amount:          31.09,
		},
	}

	for _, t := range transactions {
		suite.Require().Contains(res, t)
	}
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsTestSuite))
}
