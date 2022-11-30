package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/helper"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user entity.ProdukEntity) string
	Insert2(ctx context.Context, tx *sql.Tx, qu string) string
	// FindById(ctx context.Context, tx *sql.Tx, username int) []entity.ProdukEntity
	FindAll(ctx context.Context, tx *sql.Tx, where string) []entity.ProdukEntity
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user entity.ProdukEntity) string {

	SQL := fmt.Sprintf("insert into produks (product_id,product_name,sub_category,brand,price,status,created_by,created_time,updated_by) values ('%s','%s','%s','%s',%d,'%s','%s','%s','%s') ", user.ProductId, user.ProductName, user.SubCategory, user.Brand, user.Price, user.Status, user.CreatedBy, user.CreatedTime, user.UpdatedBy)
	// fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)

	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}

}
func (repository *UserRepositoryImpl) Insert2(ctx context.Context, tx *sql.Tx, q string) string {

	SQL := fmt.Sprintf("insert into produks (product_id,product_name,sub_category,brand,price,status,created_by,created_time,updated_by) values %s ", q)
	// fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)

	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}

}

// func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, username int) []entity.ProdukEntity {
// 	SQL := fmt.Sprintf("select id,nama,kategori,stok,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from produks where id='%d'", username)
// 	var datas []entity.ProdukEntity
// 	var data entity.ProdukEntity
// 	row, err := tx.QueryContext(ctx, SQL)
// 	fmt.Println(err, username)
// 	helper.PanicIfError(err)
// 	// fmt.Print(row)

// 	for row.Next() {
// 		err := row.Scan(&data.RowId, &data.Nama, &data.Kategori, &data.Stok, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)
// 		helper.PanicIfError(err)
// 		datas = append(datas, data)
// 	}
// 	return datas
// }

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, where string) []entity.ProdukEntity {
	SQL := fmt.Sprintf("select id,product_id,product_name,sub_category,brand,price,status,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from produks where %s", where)
	fmt.Println(SQL)
	var datas []entity.ProdukEntity
	var data entity.ProdukEntity
	row, err := tx.QueryContext(ctx, SQL)
	// fmt.Println(err, row)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.ProductId, &data.ProductName, &data.SubCategory, &data.Brand, &data.Price, &data.Status, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)
		// fmt.Println(err)

		helper.PanicIfError(err)
		// data.CreatedTime = helper.ConvertDateTime(data.CreatedTime)
		// data.UpdatedTime = helper.ConvertDateTime(data.UpdatedTime)

		datas = append(datas, data)
	}
	return datas
}
