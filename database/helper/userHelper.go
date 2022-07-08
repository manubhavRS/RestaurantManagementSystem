package helper

import (
	"github.com/elgris/sqrl"
	_ "github.com/elgris/sqrl"
	"github.com/jmoiron/sqlx"
	"log"
	"restaurantManagementSystem/database"
	"restaurantManagementSystem/models"
	"restaurantManagementSystem/models/userModels"
	"restaurantManagementSystem/utilities"
)

func CreateUser(user userModels.AddUserModel, tx *sqlx.Tx) (string, error) {
	//language=SQL
	SQL := `INSERT INTO users(name, email, password, created_by)
			VALUES ($1, $2, $3, $4)
			RETURNING user_id;`

	var userID string
	userPassword, err := utilities.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	err = tx.Get(&userID, SQL, user.Name, user.Email, userPassword, user.CreatedBy)
	if err != nil {
		log.Printf("CreateUser: Error creating User %v", err)
		return "", err
	}
	return userID, nil
}
func SignInCredential(email string) (userModels.UserModel, error) {
	//language=SQL
	SQL := `SELECT  user_id, password 
			from users 
			WHERE email=$1
			AND archived_at IS NULL;`

	var userDetails userModels.UserModel
	err := database.Rms.Get(&userDetails, SQL, email)
	if err != nil {
		log.Printf("Error SignIn User")
		return userDetails, err
	}

	return userDetails, nil
}
func AddLocation(user userModels.AddUserModel, userID string, tx *sqlx.Tx) error {
	locations := user.Location
	psql := sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)
	insertBuilder := psql.Insert("user_location").Columns("user_id", "latitude", "longitude")
	for _, location := range locations {
		insertBuilder.Values(userID, location.Latitude, location.Longitude)
	}
	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		log.Printf("Error Adding location")
		return err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		log.Printf("Error Adding location")
		return err
	}
	return nil
}

func AddUserRole(user userModels.AddUserModel, userID string, tx *sqlx.Tx) error {
	psql := sqrl.StatementBuilder.PlaceholderFormat(sqrl.Dollar)
	insertBuilder := psql.Insert("user_roles").Columns("user_id", "name", "role")
	for _, role := range user.Role {
		insertBuilder.Values(userID, user.Name, role)
	}
	sql, args, err := insertBuilder.ToSql()
	if err != nil {
		log.Printf("Error Adding roles")
		return err
	}
	_, err = tx.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}
func FetchUserLocations(userID string) ([]models.LocationModel, error) {
	//language=SQL
	SQL := `SELECT location_id, latitude, longitude
			from user_location 
			WHERE user_id=$1
			AND archived_at IS NULL`

	var locations []models.LocationModel
	err := database.Rms.Select(&locations, SQL, userID)
	if err != nil {
		log.Printf("Error Getting Location")
		return locations, err
	}

	return locations, nil
}
func FetchUserDetails(userID string) (*userModels.UserModel, error) {
	//language=SQL
	SQL := `SELECT user_id, name, email, password, created_at
			from users 
			WHERE user_id=$1
			AND archived_at IS NULL;`

	var userDetails userModels.UserModel
	err := database.Rms.Get(&userDetails, SQL, userID)
	if err != nil {
		log.Printf("Error Getting UserDetails")
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
	//language=SQL
	SQL := `SELECT role 
			from user_roles 
			WHERE user_id=$1
			AND archived_at IS NULL`

	var userRoles userModels.UserRoleModel
	roles := make([]string, 0)
	err := database.Rms.Select(&roles, SQL, userID)
	if err != nil {
		log.Printf("Error Getting User roles")
		return userRoles, err
	}
	for _, role := range roles {
		switch role {
		case utilities.AdminRole:
			userRoles.Admin = true
		case utilities.SubAdminRole:
			userRoles.SubAdmin = true
		case utilities.UserRole:
			userRoles.User = true
		}
		log.Printf(role + " ")
	}
	return userRoles, nil
}
func FetchAllUser(userID string) ([]userModels.FetchUserModel, error) {
	//language=SQL
	SQL := `SELECT user_id, name, email, created_at
			FROM users
			WHERE archived_at IS NULL
			AND user_id <> $1;`

	var users []userModels.FetchUserModel
	errRow := database.Rms.Select(&users, SQL, userID)
	if errRow != nil {

		log.Printf("Error Getting all users")
		return users, errRow
	}
	for _, u := range users {
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
	//SQL := `SELECT u.user_id, u.name, array_agg(r.role) as role
	//		FROM users u
	//		JOIN user_roles r on u.user_id = r.user_id
	//		WHERE u.archived_at IS NULL
	//		group by u.user_id`
	////var u []userModels.FetchUserModel2
	//var users []userModels.FetchUserModel
	//rows, errRow := database.Rms.Query(SQL, userID)
	//if errRow != nil {
	//	log.Printf("Error Getting all users")
	//	return users, errRow
	//}
	//for rows.Next() {
	//	var u userModels.FetchUserModel2
	//	err := rows.Scan(&u.ID, &u.Name, &u.Role)
	//	if err != nil {
	//		return users, err
	//	}
	//
	//	for _, r := range u.Role {
	//		var usr userModels.FetchUserModel
	//		var role userModels.UserRoleModel
	//		switch r {
	//		case utilities.AdminRole:
	//			role.Admin = true
	//		case utilities.SubAdminRole:
	//			role.SubAdmin = true
	//		case utilities.UserRole:
	//			role.User = true
	//		}
	//		usr.Role = role
	//		usr.ID = u.ID
	//		usr.Email = u.Email
	//		users = append(users, usr)
	//	}

	//	}

	return users, nil
}

func FetchSpecificUser(userID string) ([]userModels.FetchUserModel, error) {
	//language=SQL
	SQL := `SELECT user_id, name, email, created_at 
			FROM users 
			WHERE created_by=$1 
			AND archived_at IS NULL;`

	users := make([]userModels.FetchUserModel, 0)
	errRow := database.Rms.Select(&users, SQL, userID)
	if errRow != nil {
		log.Printf("Error Getting Specific users")
		return users, errRow
	}

	for _, u := range users {
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
	//language=SQL
	SQL := `SELECT user_id, name, email, created_at 
			FROM users 
			WHERE created_by=$1 
			AND archived_at IS NULL;`

	users := make([]userModels.FetchUserModel, 0)
	errRow := database.Rms.Select(&users, SQL, userID)
	if errRow != nil {
		log.Printf("Error Getting Subadmins")
		return users, errRow
	}

	for _, u := range users {
		roles, err := FetchUserRoles(u.ID)
		if err != nil {
			return users, err
		}
		if roles.SubAdmin && !roles.Admin {
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
	//language=SQL
	SQL := `SELECT latitude,longitude 
			FROM user_location
			WHERE location_id=$1 
			AND archived_at IS NULL;`

	var location models.LocationModel
	err := database.Rms.Select(&location, SQL, locationID)

	if err != nil {
		log.Printf("Error Getting locations")
		return location, err
	}
	return location, nil
}
func AddPrivilege(user userModels.AddPrivilegeModel) error {
	//language=SQL
	SQL := `INSERT INTO user_roles(user_id, name, role)
			VALUES ($1, $2, $3)
			RETURNING user_id;`

	_, err := database.Rms.Exec(SQL, user.ID, user.Role)
	if err != nil {
		log.Printf("Error adding privilege")
		return err
	}
	return nil
}
