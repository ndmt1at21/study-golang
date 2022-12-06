package queries

import (
	"testing"
	"unittest/models"

	"github.com/stretchr/testify/require"
)

func TestCreateUserShouldSuccess(t *testing.T) {
	type CreateUserSuccessTestCase struct {
		input          models.CreateUserData
		expectedResult models.User
	}

	testCases := []CreateUserSuccessTestCase{
		{
			input: models.CreateUserData{
				Email:    "john_doe@gmail.com",
				Name:     "John Doe",
				Password: "password",
			},
			expectedResult: models.User{
				Email:    "john_doe@gmail.com",
				Name:     "John Doe",
				Password: "password",
			},
		},
	}

	for _, val := range testCases {
		user, err := queries.CreateUser(val.input)

		require.NoError(t, err)
		require.NotEmpty(t, user)

		require.Equal(t, val.expectedResult.Email, user.Email)
		require.Equal(t, val.expectedResult.Name, user.Name)
		require.Equal(t, val.expectedResult.Password, user.Password)
	}
}
