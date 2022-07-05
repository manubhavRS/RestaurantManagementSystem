package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID string `db:"user_id" json:"user_id"`
	jwt.StandardClaims
}
