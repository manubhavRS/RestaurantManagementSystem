package helper

import (
	"log"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/models"
	"restaurantManagementSystem/models/userModels"
	"restaurantManagementSystem/utilities"
)

func CreateUser(user userModels.AddUserModel) (string, error) {
	SQL := `INSERT INTO users(name, email, password, created_by)
			VALUES ($1, $2, $3, $4)
			RETURNING user_id;`

	//log.Printf(user.Name)
	var userID string

	userPassword, err := utilities.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	err = database.Rms.QueryRow(SQL, user.Name, user.Email, interface{}(userPassword), user.CreatedBy).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}
func SignInCredential(email string) (userModels.UserModel, error) {
	SQL := `SELECT  user_id, password 
			from users 
			WHERE email=$1
			AND archived_at IS NULL;`

	var userDetails userModels.UserModel
	err := database.Rms.QueryRow(SQL, email).Scan(&userDetails.ID, &userDetails.Password)
	if err != nil {
		return userDetails, err
	}

	return userDetails, nil
}
func AddLocation(user userModels.AddUserModel, userID string) error {
	locations := user.Location

	for _, location := range locations {
		SQL := `INSERT INTO user_location(user_id, latitude, longitude) 
			VALUES($1,$2,$3) 
			RETURNING location_id;`
		err := database.Rms.QueryRow(SQL, userID, location.Latitude, location.Longitude).Scan(&location.LocationID)
		if err != nil {
			return err
		}

		log.Printf("LocationID:" + location.LocationID)
	}
	return nil
}

func AddUserRole(user userModels.AddUserModel, userID string) error {
	SQL := `INSERT INTO user_roles(user_id, name, role) 
			VALUES($1,$2,$3);`

	for _, role := range user.Role {
		_, err := database.Rms.Query(SQL, userID, user.Name, role)
		if err != nil {
			return err
		}
	}
	return nil
}
func FetchUserLocations(userID string) ([]models.LocationModel, error) {
	SQL := `SELECT location_id, latitude, longitude
			from user_location 
			WHERE user_id=$1
			AND archived_at IS NULL`

	var locations []models.LocationModel
	err := database.Rms.Select(&locations, SQL, userID)
	if err != nil {
		return locations, err
	}

	return locations, nil
}
func FetchUserDetails(userID string) (*userModels.UserModel, error) {
	SQL := `SELECT user_id, name, email, password, created_at
			from users 
			WHERE user_id=$1
			AND archived_at IS NULL;`

	var userDetails userModels.UserModel

	err := database.Rms.QueryRow(SQL, userID).Scan(&userDetails.ID, &userDetails.Name, &userDetails.Email, &userDetails.Password, &userDetails.CreatedAt)
	if err != nil {
		return &userDetails, err
	}
	location, err := FetchUserLocations(userID)

	userDetails.Location = location
	if err != nil {
		return &userDetails, err
	}

	roles, err := FetchUserRoles(userID)
	userDetails.Role = roles
	if err != nil {
		return &userDetails, err
	}

	return &userDetails, nil
}
func FetchUserRoles(userID string) (userModels.UserRoleModel, error) {
	SQL := `SELECT role 
			from user_roles 
			WHERE user_id=$1
			AND archived_at IS NULL`

	var userRoles userModels.UserRoleModel
	rows, err := database.Rms.Queryx(SQL, userID)
	if err != nil {
		return userRoles, err
	}

	for rows.Next() {
		var role string
		rows.Scan(&role)
		switch role {
		case "admin":
			userRoles.Admin = true
		case "sub-admin":
			userRoles.SubAdmin = true
		case "user":
			userRoles.User = true
		}
	}
	return userRoles, nil
}
func FetchAllUser() ([]userModels.FetchUserModel, error) {
	SQL := `SELECT user_id, name, email, created_at
			FROM users 
			WHERE archived_at IS NULL;`

	var users []userModels.FetchUserModel
	rows, errRow := database.Rms.Queryx(SQL)

	if errRow != nil {
		return users, errRow
	}

	for rows.Next() {
		var u userModels.FetchUserModel
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

		roles, err := FetchUserRoles(u.ID)
		if err != nil {
			return users, err
		}
		u.Role = roles

		locations, err := FetchUserLocations(u.ID)
		if err != nil {
			return users, err
		}
		u.Location = locations

		users = append(users, u)
	}

	return users, nil
}
func FetchSpecificUser(userID string) ([]userModels.FetchUserModel, error) {
	SQL := `SELECT user_id, name, email, created_at 
			FROM users 
			WHERE created_by=$1 
			AND archived_at IS NULL;`

	users := make([]userModels.FetchUserModel, 0)
	rows, errRow := database.Rms.Queryx(SQL, userID)
	if errRow != nil {
		return users, errRow
	}

	for rows.Next() {
		var u userModels.FetchUserModel
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

		roles, err := FetchUserRoles(u.ID)
		if err != nil {
			return users, err
		}
		u.Role = roles

		locations, err := FetchUserLocations(u.ID)
		if err != nil {
			return users, err
		}
		u.Location = locations

		users = append(users, u)
	}
	return users, nil
}

func FetchSubadminsUser(userID string) ([]userModels.FetchUserModel, error) {
	SQL := `SELECT user_id, name, email, created_at 
			FROM users 
			WHERE created_by=$1 
			AND archived_at IS NULL;`

	users := make([]userModels.FetchUserModel, 0)
	rows, errRow := database.Rms.Queryx(SQL, userID)
	if errRow != nil {
		return users, errRow
	}

	for rows.Next() {
		var u userModels.FetchUserModel
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

		roles, err := FetchUserRoles(u.ID)
		if err != nil {
			return users, err
		}
		if roles.SubAdmin {
			u.Role = roles

			locations, err := FetchUserLocations(u.ID)
			if err != nil {
				return users, err
			}
			u.Location = locations
			users = append(users, u)
		}
	}
	return users, nil
}
func FetchLocation(locationID string) (models.LocationModel, error) {
	SQL := `SELECT latitude,longitude 
			FROM user_location
			WHERE location_id=$1 
			AND archived_at IS NULL;`

	var location models.LocationModel
	err := database.Rms.QueryRow(SQL, locationID).Scan(&location.Latitude, &location.Longitude)

	if err != nil {
		return location, err
	}
	return location, nil
}
