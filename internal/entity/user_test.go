package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserEntityCreation(t *testing.T) {
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
	for _, scenario := range []struct {
		description string
		name        string
		email       string
		password    string
		wantErr     bool
		errMessage  string
	}{
		{
			description: "should create a new user",
			name:        "Caique Ribeiro",
			email:       "caique@gmail.com",
			password:    "123456",
			wantErr:     false,
		},
		{
			description: "should return error if name is not informed",
			name:        "",
			email:       "caique@gmail.com",
			password:    "123456",
			wantErr:     true,
			errMessage:  "user name is required",
		},
		{
			description: "should return error if email is not informed",
			name:        "Caique Ribeiro",
			email:       "",
			password:    "123456",
			wantErr:     true,
			errMessage:  "email is required",
		},
		{
			description: "should return error if password is not informed",
			name:        "Caique Ribeiro",
			email:       "caique@gmail.com",
			password:    "",
			wantErr:     true,
			errMessage:  "password is required",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			user, err := NewUser(scenario.name, scenario.email, string(scenario.password))

			if scenario.wantErr {
				println(scenario.description, err.Error())
				assert.NotNil(t, err)
				assert.EqualError(t, err, scenario.errMessage)
				assert.Nil(t, user)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, user.Name, scenario.name)
				assert.Equal(t, user.Email, scenario.email)
				assert.NotNil(t, user.Password)
			}
		})
	}
}
