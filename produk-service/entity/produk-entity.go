package entity

type ProdukEntity struct {
	RowId       int    `json:"id"`
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	SubCategory string `json:"sub_category"`
	Brand       string `json:"brand"`
	Price       int    `json:"price"`
	CreatedBy   string `json:"created_by"`
	CreatedTime string `json:"created_time"`
	UpdatedBy   string `json:"updated_by"`
	UpdatedTime string `json:"updated_time"`
	Status      string `json:"status"`
}

type UpdateUserEntity struct {
	RowId    string `json:"id"`
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
	Stok     int    `json:"stok"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
