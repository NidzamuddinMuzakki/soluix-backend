package entity

type ProdukEntity struct {
	RowId       int    `json:"id"`
	Nama        string `json:"nama"`
	Kategori    string `json:"kategori"`
	Stok        int    `json:"stok"`
	CreatedBy   string `json:"created_by"`
	CreatedTime string `json:"created_time"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedTime string `json:"updated_time"`
	Username    string `json:"username"`
	Role        string `json:"role"`
}

type UpdateUserEntity struct {
	RowId    string `json:"id"`
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
	Stok     int    `json:"stok"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
