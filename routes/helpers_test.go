package routes

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/errorh"
	"github.com/arfanxn/coursefan-gofiber/app/helpers/rwh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/models"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/clarketm/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// ActingAs will sign in a user with the given input.
// Function parameters descriptions
//   - input (the login creadentials)
//   - httpRequest (this will be used as login client, so this parameter is required)
//   - httpResponse (this parameter is optional, it can be passed by nil value)
func ActingAs(
	t *testing.T,
	input requests.AuthLogin,
	httpRequest *http.Request,
	httpResponse *http.Response,
) (
	result struct {
		input        requests.AuthLogin
		user         models.User
		httpRequest  *http.Request
		httpResponse *http.Response
	},
) {
	t.Helper()
	require := require.New(t)
	var (
		urlStr     string = "/api/users/login"
		httpMethod string = http.MethodPost
	)

	db, err := databasep.GetGormDB()
	require.Nil(err)
	userRepository := repositories.NewUserRepository(db)
	c := new(fiber.Ctx)
	user, err := userRepository.FindByEmail(c, input.Email)
	// Check if user is exists in database , if not create it
	if errorh.IsGormErrRecordNotFound(err) {
		err = nil
		user = factories.FakeUser()
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		require.Nil(err)
		user.Password = string(hashedPasswordBytes)
		affected, err := userRepository.Insert(c, &user)
		require.Nil(err)
		require.NotZero(affected)
	}
	require.Nil(err)

	// Ensure that the httpRequest parameter is not nil
	require.NotNil(httpRequest)
	urlParsed, err := url.Parse(urlStr)
	require.Nil(err)
	httpRequest.URL = urlParsed
	httpRequest.Method = httpMethod
	bodyReadCloser, err := rwh.AnyToReadCloser(input)
	require.Nil(err)
	httpRequest.Body = bodyReadCloser

	// Set header content type
	httpRequest.Header.Set("Content-Type", "application/json")
	// Do Login request
	httpResponse, err = testApp.Test(httpRequest, -1)
	require.Nil(err)
	// Expected response
	expectedResponse := resources.Response{
		Code:    fiber.StatusOK,
		Message: "Login successfully",
	}
	// Actual response
	var actualResponse resources.Response
	err = json.NewDecoder(httpResponse.Body).Decode(&actualResponse)
	require.Nil(err)
	// Assert response
	require.Equal(expectedResponse.Code, httpResponse.StatusCode)
	require.Equal(expectedResponse.Code, actualResponse.Code)
	httpRequest.Body.Close() // close body before go to

	// return the result
	result.input = input
	result.user = user
	result.httpRequest = httpRequest
	result.httpResponse = httpResponse
	return
}
