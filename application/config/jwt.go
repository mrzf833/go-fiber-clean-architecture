package config

import (
	"go-fiber-clean-architecture/application/helper/helper2"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.Claims
}

var JwtKey = helper2.GetEnv("JWT_KEY", "secret")

// set expire token = default 24 hours
var ExpireToken = time.Hour * 24
