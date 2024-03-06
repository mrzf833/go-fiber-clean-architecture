package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.Claims
}

var JwtKey = Config("JWT_KEY", "secret")

// set expire token = default 24 hours
var ExpireToken = time.Hour * 24
