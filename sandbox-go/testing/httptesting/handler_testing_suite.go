package httptesting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite

	server *httptest.Server

	Handler http.Handler
}

// Interface implementation
func (suite *HandlerTestSuite) SetupSuite() {
	suite.server = httptest.NewServer(suite.Handler)
}

func (suite *HandlerTestSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *HandlerTestSuite) TestPulic() {
	expect := Expect(suite.server.URL, suite.T())

	e := expect.GET("/public").Expect()
	e.Status(http.StatusOK).Body().Equal("this is public endpoint")
}

func (suite *HandlerTestSuite) TestAuthorizeMethodPost() {
	expect := Expect(suite.server.URL, suite.T())
	e := expect.GET("/authorize").Expect()
	o := e.Status(http.StatusMethodNotAllowed).JSON()
	o.Object().ValueEqual("message", "Method Not Allowed")
}

func (suite *HandlerTestSuite) TestAuthorizeLogin() {
	// getting token
	type Login struct {
		Username string `form:"user"`
		Password string `form:"pass"`
	}

	tests := []struct {
		name       string
		user, pass string
		code       int
	}{
		{
			name: "bad password",
			user: "foo",
			pass: "baz",
			code: http.StatusUnauthorized,
		},
		{
			name: "bad user",
			user: "fuz",
			pass: "bar",
			code: http.StatusUnauthorized,
		},
		{
			name: "ok user&password",
			user: "foo",
			pass: "bar",
			code: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		suite.T().Run(test.name, func(tt *testing.T) {
			expect := Expect(suite.server.URL, tt)

			e := expect.POST("/authorize").WithForm(Login{test.user, test.pass}).Expect()
			e.Status(test.code)
		})
	}
}

func (suite *HandlerTestSuite) TestRestrictedStatus() {
	expect := Expect(suite.server.URL, suite.T())

	// getting token
	type Login struct {
		Username string `form:"user"`
		Password string `form:"pass"`
	}
	r := expect.POST("/authorize").WithForm(Login{"foo", "bar"}).
		Expect().
		Status(http.StatusOK).JSON().Object()

	tests := []struct {
		name       string
		hKey, hVal string
		code       int
	}{
		{
			name: "Bad Header Token",
			hKey: "foo",
			hVal: "baz",
			code: http.StatusBadRequest,
		},
		{
			name: "Bad Token",
			hKey: "Authorization",
			hVal: "Bearer <bad token>",
			code: http.StatusUnauthorized,
		},
		{
			name: "Good token",
			hKey: "Authorization",
			hVal: "Bearer " + r.Value("token").String().Raw(),
			code: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		suite.T().Run(test.name, func(tt *testing.T) {
			expect := Expect(suite.server.URL, tt)

			e := expect.POST("/restricted").WithHeader(test.hKey, test.hVal).Expect()
			e.Status(test.code)
		})
	}
}

func (suite *HandlerTestSuite) TestRestrictedContent() {
	expect := Expect(suite.server.URL, suite.T())

	// getting token
	type Login struct {
		Username string `form:"user"`
		Password string `form:"pass"`
	}

	r := expect.POST("/authorize").WithForm(Login{"foo", "bar"}).
		Expect().
		Status(http.StatusOK).JSON().Object()

	e := expect.POST("/restricted").
		WithHeader("Authorization", "Bearer "+r.Value("token").String().Raw()).
		Expect()
	e.Status(http.StatusOK).Body().Equal("Welcome, foo\n")
}

// wrapper for httpexpect, so it fail with correct test id.
func Expect(url string, t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  url,
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: nil,
	})
}
