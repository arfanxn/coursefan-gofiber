package routes

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/jsonh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/app/providers/databasep"
	"github.com/arfanxn/coursefan-gofiber/app/repositories"
	"github.com/arfanxn/coursefan-gofiber/database/factories"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/clarketm/json"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAuthRouter(t *testing.T) {
	require := require.New(t)

	// Prepare required dependencies
	c := new(fiber.Ctx)
	db, err := databasep.GetGormDB()
	require.Nil(err)
	userRepository := repositories.NewUserRepository(db)

	t.Run("Register", func(t *testing.T) {
		var input requests.AuthRegister
		input.Name = faker.Name()
		input.Email = faker.Email()
		input.Password = "111222"
		input.ConfirmPassword = input.Password

		urlStr, httpMethod := "/api/users/register", fiber.MethodPost
		httpRequest := httptest.NewRequest(
			httpMethod,
			urlStr,
			bytes.NewReader(jsonh.MustMarshal(input)),
		)
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse, err := testApp.Test(httpRequest, -1)
		require.Nil(err)
		defer httpRequest.Body.Close()

		expectedResponse := resources.Response{
			Code:    fiber.StatusCreated,
			Message: "Register successfully",
		}
		var actualResponse resources.Response
		err = json.NewDecoder(httpResponse.Body).Decode(&actualResponse)
		require.Nil(err)

		require.Equal(expectedResponse.Code, httpResponse.StatusCode)
		require.Equal(expectedResponse.Code, actualResponse.Code)
	})

	t.Run("Login", func(t *testing.T) {
		password := "111222"
		hashedPassword := "$2a$10$1sGm.uAbtb6h9HkZv1/5S.IFesDq7GOJx0gjXAhGltA3hFssCs/kO"
		user := factories.User()
		user.Password = hashedPassword
		affected, err := userRepository.Insert(c, &user)
		require.Nil(err)
		require.NotZero(affected)

		input := requests.AuthLogin{
			Email:    user.Email,
			Password: password,
		}
		urlStr, httpMethod := "/api/users/login", fiber.MethodPost
		httpRequest := httptest.NewRequest(
			httpMethod,
			urlStr,
			bytes.NewReader(jsonh.MustMarshal(input)),
		)
		httpRequest.Header.Set("Content-Type", "application/json")
		httpResponse, err := testApp.Test(httpRequest, -1)
		require.Nil(err)
		defer httpRequest.Body.Close()

		expectedResponse := resources.Response{
			Code:    fiber.StatusOK,
			Message: "Login successfully",
		}
		var actualResponse resources.Response
		err = json.NewDecoder(httpResponse.Body).Decode(&actualResponse)
		require.Nil(err)

		require.Equal(expectedResponse.Code, httpResponse.StatusCode)
		require.Equal(expectedResponse.Code, actualResponse.Code)
	})
}
