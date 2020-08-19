package ziggy

import (
	"context"
	"testing"
	"time"

	"github.com/nickcorin/investec/mock"
	"github.com/stretchr/testify/suite"
)

type TransactionsTestSuite struct {
	suite.Suite
	client Client
	server *mock.Server
}

func (suite *TransactionsTestSuite) SetupSuite() {
	suite.server = mock.NewServer()
	suite.client = NewForTesting(suite.T(), suite.server.URL, nil)
}

func (suite *TransactionsTestSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *TransactionsTestSuite) TestClient_GetAccountTransactions() {
	res, err := suite.client.GetAccountTransactions(context.TODO(),
		&TransactionsRequest{AccountID: "123456789"})
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	transactions := []Transaction{
		{
			AccountID:       "172878438321553632224",
			Type:            TransactionTypeDebit,
			Status:          TransactionStatusPosted,
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
			Type:            TransactionTypeCredit,
			Status:          TransactionStatusPosted,
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
		suite.Require().Contains(res.Data.Transactions, t)
	}
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsTestSuite))
}
