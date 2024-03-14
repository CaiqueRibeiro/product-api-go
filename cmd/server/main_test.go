package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CaiqueRibeiro/product-api/configs"
	"github.com/stretchr/testify/assert"
)

// to run: go test -v . -run '^[^_gen].*$'
func TestApp(t *testing.T) {
	configs, err := configs.LoadConfig(".env")

	if err != nil {
		panic(err)
	}

	ts := httptest.NewServer(BuildServer(configs.JWTSecret, configs.JWTExpiresIn))
	defer ts.Close()

	// Get the JWT token
	// Define the data
	data := map[string]string{
		"email":    "zezinho@gmail.com",
		"password": "123456",
	}

	// Encode the data to JSON
	payload, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	res, _ := http.Post(ts.URL+"/users/generate_token", "application/json", bytes.NewBuffer(payload))
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	var response map[string]string
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatal(err)
	}

	// Now you can access the access_token
	token := response["access_token"]

	// Create a new request
	req, err := http.NewRequest("GET", ts.URL+"/products", nil)
	assert.Nil(t, err)

	req.Header.Add("Authorization", token)

	res, err = ts.Client().Do(req)
	productsBody, _ := io.ReadAll(res.Body)
	defer res.Body.Close()

	var products []struct {
		ID        string  `json:"id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		CreatedAt string  `json:"created_at"`
	}

	if err := json.Unmarshal(productsBody, &products); err != nil {
		t.Fatal(err)
	}

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, len(products), 2)
}
