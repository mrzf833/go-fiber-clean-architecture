package config

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.Claims
}

var JwtKey = Config("JWT_KEY", "secret")

var ExpireToken = time.Hour * 24
