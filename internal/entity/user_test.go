package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		user, err := NewUser(
			"Caique Ribeiro",
			"caique@gmail.com",
			"123456",
		)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.Equal(t, user.Name, "Caique Ribeiro")
		assert.Equal(t, user.Email, "caique@gmail.com")
	})

	t.Run("should validate with success", func(t *testing.T) {
		user := &User{
			Name:     "Caique Ribeiro",
			Email:    "caique@gmail.com",
			Password: "123456",
		}

		err := user.Validate()
		assert.NoError(t, err)
	})

	t.Run("should validate password with success", func(t *testing.T) {
		user, _ := NewUser(
			"Caique Ribeiro",
			"caique@gmail.com",
			"123456",
		)

		assert.True(t, user.ValidatePassword("123456"))
	})

	t.Run("should validate password with failure", func(t *testing.T) {
		user, _ := NewUser(
			"Caique Ribeiro",
			"caique@gmail.com",
			"123456",
		)

		assert.False(t, user.ValidatePassword("123"))
	})

	type errorTestCases struct {
		description   string
		input         *User
		expectedError string
	}

	for _, scenario := range []errorTestCases{
		{
			description: "should return error if name is not informed",
			input: &User{
				Name:     "",
				Email:    "caique@gmail.com",
				Password: "123456",
			},
			expectedError: "user name is required",
		},
		{
			description: "should return error if email is not informed",
			input: &User{
				Name:     "Caique Ribeiro",
				Email:    "",
				Password: "123456",
			},
			expectedError: "email is required",
		},
		{
			description: "should return error if password is not informed",
			input: &User{
				Name:     "Caique Ribeiro",
				Email:    "caique@gmail.com",
				Password: "",
			},
			expectedError: "password is required",
		},
		{
			description: "should return error if password is not informed",
			input: &User{
				Name:     "Caique Ribeiro",
				Email:    "caique@gmail.com",
				Password: "",
			},
			expectedError: "password is required",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			err := scenario.input.Validate()
			assert.EqualError(t, err, scenario.expectedError)
		})
	}
}
