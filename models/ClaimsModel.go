package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID string `db:"user_id" json:"userID"`
	jwt.StandardClaims
}
