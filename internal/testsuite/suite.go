package testsuite

import (
	"time"

	db "github.com/star-table/db/v4"
	"github.com/stretchr/testify/suite"
)

const TimeZone = "Canada/Eastern"

var defaultTimeLocation, _ = time.LoadLocation(TimeZone)

type Helper interface {
	Session() db.Session

	Adapter() string

	TearUp() error
	TearDown() error
}

type Suite struct {
	suite.Suite

	Helper
}

func (s *Suite) AfterTest(suiteName, testName string) {
	err := s.TearDown()
	s.NoError(err)
}

func (s *Suite) BeforeTest(suiteName, testName string) {
	err := s.TearUp()
	s.NoError(err)
}
