package investec

import (
	"context"
	"net/http"
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

func (suite *TransactionsTestSuite) SetupTest(h http.HandlerFunc) {
	suite.server = mock.NewServer(h)
	suite.client = NewClient(WithBaseURL(suite.server.URL))
}

func (suite *TransactionsTestSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *TransactionsTestSuite) TestClient_GetAccountTransactions() {
	suite.SetupTest(suite.server.GetAccountTransactions)

	res, err := suite.client.GetAccountTransactions(context.TODO(),
		&TransactionsRequest{AccountID: "123ABC"})

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
			TransactionDate: time.Date(2020, 6, 10, 0, 0, 0, 0, time.UTC),
			Amount:          535,
		},
		{
			AccountID:   "172878438321553632224",
			Type:        TransactionTypeCredit,
			Status:      TransactionStatusPosted,
			Description: "CREDIT INTEREST",
			CardNumber:  "",
			PostingDate: time.Date(2020, 06, 11, 0, 0, 0, 0, time.UTC),
			ValueDate:   time.Date(2020, 06, 10, 0, 0, 0, 0, time.UTC),
			ActionDate:  time.Date(2020, 06, 18, 0, 0, 0, 0, time.UTC),
			Amount:      31.09,
		},
	}

	for _, t := range transactions {
		suite.Require().Contains(res.Data.Transactions, t)
	}
}

func TestTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsTestSuite))
}
