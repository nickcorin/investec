package config_test

import (
	"os"
	"testing"

	"github.com/nickcorin/ziggy/pkg/config"

	"github.com/stretchr/testify/suite"
)

type configTestSuite struct {
	suite.Suite
	filename string
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(configTestSuite))
}

func (suite *configTestSuite) SetupSuite() {
	suite.filename = ".test.conf"
}

func (suite *configTestSuite) TestClean() {
	path, err := config.DefaultPath(suite.filename)
	suite.Require().NoError(err)
	suite.Require().NotEmpty(path)

	_, err = os.Stat(path)
	suite.Require().True(os.IsNotExist(err))

	err = config.WriteDefault(suite.filename, true)
	suite.Require().NoError(err)

	_, err = os.Stat(path)
	suite.Require().NoError(err)

	err = config.Clean(suite.filename)
	suite.Require().NoError(err)

	_, err = os.Stat(path)
	suite.Require().True(os.IsNotExist(err))
}

func (suite *configTestSuite) TestLoad() {
	defer config.Clean(suite.filename)

	conf, err := config.Load(suite.filename)
	suite.Require().Error(err)
	suite.Require().Nil(conf)

	err = config.WriteDefault(suite.filename, false)
	suite.Require().NoError(err)

	conf, err = config.Load(suite.filename)
	suite.Require().NoError(err)
	suite.Require().NotNil(conf)
}

func (suite *configTestSuite) TestWriteDefault() {
	defer config.Clean(suite.filename)

	err := config.WriteDefault(suite.filename, false)
	suite.Require().NoError(err)

	err = config.WriteDefault(suite.filename, false)
	suite.Require().NoError(err)

	err = config.WriteDefault(suite.filename, true)
	suite.Require().NoError(err)
}
