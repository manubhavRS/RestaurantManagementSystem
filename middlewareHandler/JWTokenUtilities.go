package middlewareHandler

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"restaurantManagementSystem/database/helper"
	"restaurantManagementSystem/models/userModels"
	"time"
)

const secretkey string = "SuperSecretKey"
const ContextUserKey string = "user"
const ContextRefreshToken string = "refreshToken"

func UserFromContext(ctx context.Context) *userModels.UserModel {
	return ctx.Value(ContextUserKey).(*userModels.UserModel)
}
func TokenFromContext(ctx context.Context) string {
	return ctx.Value(ContextRefreshToken).(string)
}

func GenerateJWT(user userModels.UserModel) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func JWTAuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userID := fmt.Sprint(claims["user_id"])
		users, err := helper.FetchUserDetails(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		roles, err := helper.FetchUserRoles(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users.Role = roles

		locations, err := helper.FetchUserLocations(userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		users.Location = locations
		refreshToken, err := GenerateJWT(*users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Refresh Token: " + refreshToken)
		ctx := context.WithValue(r.Context(), ContextUserKey, users)
		log.Printf("JWT Token verified...")
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["RefreshToken"] = refreshToken
		jsonResponse, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
