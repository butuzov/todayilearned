package testcontainer

import "github.com/stretchr/testify/suite"

type RepositoryTestSuite struct {
	suite.Suite

	Repo Repository
}

// Defining Test Suite

// Interface implementation
func (suite *RepositoryTestSuite) SetupSuite()                  { println("SetupSuite") }
func (suite *RepositoryTestSuite) TearDownTest()                { println("TearDownTest") }
func (suite *RepositoryTestSuite) BeforeTest(name, test string) { println("Before", name, test) }
func (suite *RepositoryTestSuite) AfterTest(name, test string)  { println("After", name, test) }
func (suite *RepositoryTestSuite) SetupTest()                   { println("SetupTest") }
func (suite *RepositoryTestSuite) TearDownSuite()               { println("TearDownSuite") }

func (suite *RepositoryTestSuite) TestSum() {
	suite.Equal(3, suite.Repo.Sum(1, 2))
}
