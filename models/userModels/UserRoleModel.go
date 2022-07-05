package userModels

type UserRoleModel struct {
	Admin    bool `db:"admin" json:"admin"`
	SubAdmin bool `db:"sub-admin" json:"sub-admin"`
	User     bool `db:"user" json:"user"`
}
