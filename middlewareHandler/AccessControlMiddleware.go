package middlewareHandler

import (
	"encoding/json"
	"net/http"
	"restaurantManagementSystem/models/userModels"
	"restaurantManagementSystem/utilities"
)

func AccessControlMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var signedUser *userModels.UserModel
		signedUser = UserFromContext(r.Context())
		if !signedUser.Role.Admin && !signedUser.Role.SubAdmin {
			if r.RequestURI == utilities.Path2 && r.RequestURI == utilities.Path1 && r.RequestURI == utilities.Path3 && r.RequestURI == utilities.Path4 && r.RequestURI == utilities.Path6 && r.RequestURI == utilities.Path7 && r.RequestURI == utilities.Path9 && r.RequestURI == utilities.Path11 && r.RequestURI == utilities.Path13 && r.RequestURI == utilities.Path14 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else if !signedUser.Role.Admin && signedUser.Role.SubAdmin {
			if r.RequestURI == utilities.Path2 && r.RequestURI == utilities.Path4 && r.RequestURI == utilities.Path8 && r.RequestURI == utilities.Path11 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else if r.RequestURI == utilities.Path1 {
				var addUser userModels.AddUserModel
				addErr := json.NewDecoder(r.Body).Decode(&addUser)
				if addErr != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				for _, role := range addUser.Role {
					if role == utilities.AdminRole || role == utilities.SubAdminRole {
						w.WriteHeader(http.StatusUnauthorized)
						return
					}
				}
			}
		}
		handler.ServeHTTP(w, r)
	})
}
