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

func TestGuestRegisterSuccess(t *testing.T) {
	DeleteUser(t, "test")

	jsonData, err := json.Marshal(map[string]string{
		"username": "test",
		"name":     "test",
		"password": "Password",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	url := "http://localhost:4000/api/auth/register"
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

	_, data := body["data"]
	_, message := body["message"]

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.True(t, data)
	assert.True(t, message)
}

func TestGuestRegisterFail(t *testing.T) {
	DeleteUser(t, "test")

	jsonData, err := json.Marshal(map[string]string{
		"username": "",
		"name":     "",
		"password": "",
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	url := "http://localhost:4000/api/auth/register"
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
	assert.True(t, errors)
	assert.True(t, message)
}
