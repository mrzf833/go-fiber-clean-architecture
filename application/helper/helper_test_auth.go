package helper

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

var Token = ""

type LoginResponse struct {
	Expire string
	Token string
	Username string
}

func LoginAuth(appRun *fiber.App, ch chan string) {
	body := strings.NewReader(`{"username":"john","password":"doe"}`)
	req := httptest.NewRequest(fiber.MethodPost, "/api/login", body)
	req.Header.Set("Content-Type", "application/json")

	// panggil handler
	res, err := appRun.Test(req)

	if err != nil{
		panic("Failed to login: " + err.Error())
	}

	if res.StatusCode == http.StatusOK {
		//Token = res.Header.Get("Authorization")
		bytes, _ := io.ReadAll(res.Body)
		var dataJson *LoginResponse
		json.Unmarshal(bytes, &dataJson)
		Token = dataJson.Token
		ch <- Token
	}else {
		panic("Failed to login")
	}
}
