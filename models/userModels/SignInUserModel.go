package userModels

type SignInUserModel struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
