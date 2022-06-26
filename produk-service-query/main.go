package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/app"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/rabbitmq"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/entity"
	// "github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/repository"
	// "github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/service"
)

var (
	rc rabbitmq.RabbitClient
	// db             *sql.DB                   = app.Init()
	// UserRepository repository.UserRepository = repository.NewUserRepository()
	// UserService    service.UserService       = service.NewUserService(UserRepository, db)
	ProdukSearchController controller.ProdukSearchController = controller.NewProdukSearchController()
)

func main() {
	// defer db.Close()
	var rc rabbitmq.RabbitClient
	go rc.Consume("insert-queue", InsertConsume)
	go rc.Consume("update-queue", UpdateConsume)
	go rc.Consume("delete-queue", DeleteConsume)
	r := app.InitRouter(ProdukSearchController)
	r.Start(":9003")
}

type funcName struct {
	Nama string `json:"nama"`
}

// func InsertElasticSearch(produk entity.ProdukEntity, es *elastics.PostStorage) error {
// 	nor := time.Now()
// 	doc := storages.Post{
// 		ID:        strconv.Itoa(produk.RowId),
// 		Nama:      produk.Nama,
// 		Kategori:  produk.Kategori,
// 		Stok:      produk.Stok,
// 		CreatedAt: &nor,
// 	}

// 	errs := es.Insert(context.Background(), doc)
// 	if errs != nil {
// 		panic(errs)
// 	}
// 	return nil
// }
func InsertConsume(nid interface{}) error {
	m := nid.(map[string]interface{})

	var produk entity.ProdukEntity
	if rowId, ok := m["id"].(float64); ok {
		produk.RowId = int(rowId)
	}
	fmt.Println(nid, "coba", m["id"], reflect.TypeOf(m["id"]), "hayy", produk.RowId)
	if nama, ok := m["nama"].(string); ok {
		produk.Nama = nama
	}
	if kategori, ok := m["kategori"].(string); ok {
		produk.Kategori = kategori
	}
	if stok, ok := m["stok"].(float64); ok {
		produk.Stok = int(stok)
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

	client := &http.Client{}
	getbase := fmt.Sprintf("http://%s", os.Getenv("ELASTICSEARCH_URL"))
	dataRequest := fmt.Sprintf(`{"id":%d,"nama":"%s","kategori":"%s","stok":%d}`, produk.RowId, produk.Nama, produk.Kategori, produk.Stok)
	requestBody := strings.NewReader(dataRequest)
	req, err := http.NewRequest("PUT", getbase+"/produk-v1/_doc/"+strconv.Itoa(produk.RowId), requestBody)
	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, errs := client.Do(req)
	if errs != nil {
		log.Println(errs)
	} else {
		res.Body.Close()
	}
	return nil
}

func UpdateConsume(nid interface{}) error {
	fmt.Println(nid)
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
