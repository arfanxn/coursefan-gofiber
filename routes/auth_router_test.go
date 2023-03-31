package routes

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/arfanxn/coursefan-gofiber/app/helpers/jsonh"
	"github.com/arfanxn/coursefan-gofiber/app/http/requests"
	"github.com/arfanxn/coursefan-gofiber/resources"
	"github.com/clarketm/json"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestAuthRouter(t *testing.T) {
	require := require.New(t)

	t.Run("Register", func(t *testing.T) {
		var input requests.AuthRegister
		input.Name = faker.Name()
		input.Email = faker.Email()
		input.Password = "111222"
		input.ConfirmPassword = input.Password

		urlStr := "/api/users/register"
		httpRequest := httptest.NewRequest(
			fiber.MethodPost,
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
	})
}
