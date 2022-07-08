package userModels

type AddPrivilegeModel struct {
	ID   string `db:"user_id" json:"userID"`
	Role string `db:"role" json:"role"`
}
