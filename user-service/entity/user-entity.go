package entity

type UserEntity struct {
	RowId       int    `json:"rowId"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	CreatedBy   string `json:"created_by"`
	CreatedTime string `json:"created_time"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedTime string `json:"updated_time"`
}

type UpdateUserEntity struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
