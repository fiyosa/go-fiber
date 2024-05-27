package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGuestLoginSuccess(t *testing.T) {
	DeleteUser(t, "test")
	CreateUser(t)

	jsonData, err := json.Marshal(map[string]string{
		"username": "test",
		"password": "Password",
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	url := "http://localhost:4000/api/auth/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}

	var body map[string]any
	err = json.Unmarshal(resBody, &body)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, token := body["token"]

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.True(t, token)
}

func TestGuestLoginFail(t *testing.T) {
	DeleteUser(t, "test")
	CreateUser(t)

	jsonData, err := json.Marshal(map[string]string{
		"username": "",
		"password": "",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	url := "http://localhost:4000/api/auth/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
	}

	var body map[string]any
	err = json.Unmarshal(resBody, &body)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, message := body["message"]
	_, errors := body["errors"]

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	assert.True(t, message)
	assert.True(t, errors)
}
