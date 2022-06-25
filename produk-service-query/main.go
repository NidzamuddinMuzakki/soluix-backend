package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/app"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/rabbitmq"

	// "github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/entity"
	// "github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/repository"
	// "github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/service"
)

var (
	rc rabbitmq.RabbitClient
	// db             *sql.DB                   = app.Init()
	// UserRepository repository.UserRepository = repository.NewUserRepository()
	// UserService    service.UserService       = service.NewUserService(UserRepository, db)
	// UserController controller.UserController = controller.NewUserController(UserService, rc)
)

func main() {
	// defer db.Close()
	var rc rabbitmq.RabbitClient
	go rc.Consume("insert-queue", InsertConsume)
	go rc.Consume("update-queue", UpdateConsume)
	go rc.Consume("delete-queue", DeleteConsume)

	r := app.InitRouter()
	r.Start(":9003")
}

type funcName struct {
	Nama string `json:"nama"`
}

func InsertConsume(nid interface{}) error {
	m := nid.(map[string]interface{})
	var produk entity.ProdukEntity
	if rowId, ok := m["id"].(int); ok {
		produk.RowId = rowId
	}
	if nama, ok := m["nama"].(string); ok {
		produk.Nama = nama
	}
	if kategori, ok := m["kategori"].(string); ok {
		produk.Kategori = kategori
	}
	if stok, ok := m["stok"].(int); ok {
		produk.Stok = stok
	}
	if created_by, ok := m["created_by"].(string); ok {
		produk.CreatedBy = created_by
	}
	if created_time, ok := m["created_time"].(string); ok {
		produk.CreatedTime = created_time
	}
	if updated_by, ok := m["updated_by"].(string); ok {
		produk.UpdatedBy = updated_by
	}
	if updated_time, ok := m["updated_time"].(string); ok {
		produk.UpdatedTime = updated_time
	}
	baseURL := fmt.Sprintf("http://localhost:9000/produk-v1/_doc/%d", produk.RowId)
	dataRequest := fmt.Sprintf(`{"id":%d,"nama":"%s","kategori":"%s","stok":%d}`, produk.RowId, produk.Nama, produk.Kategori, produk.Stok)
	requestBody := strings.NewReader(dataRequest)
	url := baseURL
	fmt.Println(url, "cek url")
	res, err := http.Post(url, "application/json", requestBody)

	fmt.Println(err, res, "nidzazazaza")
	dec := json.NewDecoder(res.Body)
	var p entity.WebResponseListAndDetail
	// fmt.Println(dec, res, res.Body, p, "nidzam")
	err = dec.Decode(&p)
	fmt.Println(produk)

	return nil
}

func UpdateConsume(nid interface{}) error {
	// fmt.Println(nid)
	m := nid.(map[string]interface{})
	var produk entity.ProdukEntity
	if rowId, ok := m["id"].(int); ok {
		produk.RowId = rowId
	}
	if nama, ok := m["nama"].(string); ok {
		produk.Nama = nama
	}
	if kategori, ok := m["kategori"].(string); ok {
		produk.Kategori = kategori
	}
	if stok, ok := m["stok"].(int); ok {
		produk.Stok = stok
	}
	if created_by, ok := m["created_by"].(string); ok {
		produk.CreatedBy = created_by
	}
	if created_time, ok := m["created_time"].(string); ok {
		produk.CreatedTime = created_time
	}
	if updated_by, ok := m["updated_by"].(string); ok {
		produk.UpdatedBy = updated_by
	}
	if updated_time, ok := m["updated_time"].(string); ok {
		produk.UpdatedTime = updated_time
	}
	fmt.Println(produk)
	return nil
}

func DeleteConsume(nid interface{}) error {
	// fmt.Println(nid)
	m := nid.(map[string]interface{})
	var produk entity.ProdukEntity

	if rowId, ok := m["id"].(int); ok {
		produk.RowId = rowId
	}
	fmt.Println(produk)
	return nil
}
