//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func createJWTToken() string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenStr, err := token.SignedString([]byte("mohTheBatman"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenStr)
	return tokenStr
}

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer "+createJWTToken()).
			SetBody(`"slug":"/go","body":"Go is Awesome!!","author":"Mo"`).
			Post("http://localhost:8080/api/v1/comment")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())

	})

	t.Run("cannot post comment without JWT", func(t *testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`"slug":"/go","body":"Go is Awesome!!","author":"Mo"`).
			Post("http://localhost:8080/api/v1/comment")

		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())

	})
}
