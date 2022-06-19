package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/helper"
)

type UserRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, user entity.ProdukEntity) string
	Update(ctx context.Context, tx *sql.Tx, user entity.ProdukEntity) string
	Delete(ctx context.Context, tx *sql.Tx, rowId int) string
	FindById(ctx context.Context, tx *sql.Tx, username int) []entity.ProdukEntity
	FindAll(ctx context.Context, tx *sql.Tx, where string) []entity.ProdukEntity
}

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user entity.ProdukEntity) string {
	cuyNow := helper.TimePlus7(time.Now())
	user.CreatedTime = cuyNow
	user.CreatedBy = user.Username

	SQL := fmt.Sprintf("insert into produks (nama,kategori,stok,created_by,created_time,updated_by) values ('%s','%s',%d,'%s','%s','%s') ", user.Nama, user.Kategori, user.Stok, user.CreatedBy, user.CreatedTime, user.UpdatedBy)
	fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	fmt.Println(errs, rows)
	if rows < 1 {
		helper.PanicIfError(errs)
	}
	fmt.Println(rows)
	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}

}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.ProdukEntity) string {
	fmt.Println("hay")
	cuyNow := helper.TimePlus7(time.Now())
	user.UpdatedTime = cuyNow
	user.UpdatedBy = user.Username
	SQL := fmt.Sprintf("update produks set nama='%s',kategori='%s',stok=%d,updated_time='%s',updated_by='%s' where id=%d", user.Nama, user.Kategori, user.Stok, user.UpdatedTime, user.UpdatedBy, user.RowId)
	fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)
	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, rowId int) string {
	SQL := fmt.Sprintf("delete from produks where id=%d", rowId)
	fmt.Println(SQL)
	row, err := tx.ExecContext(ctx, SQL)
	fmt.Println(err, row)
	helper.PanicIfError(err)
	rows, errs := row.RowsAffected()
	helper.PanicIfError(errs)
	if rows > 0 {
		return "berhasil"
	} else {
		return "gagal"
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, username int) []entity.ProdukEntity {
	SQL := fmt.Sprintf("select id,nama,kategori,stok,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from produks where id='%d'", username)
	var datas []entity.ProdukEntity
	var data entity.ProdukEntity
	row, err := tx.QueryContext(ctx, SQL)
	fmt.Println(err, username)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.Nama, &data.Kategori, &data.Stok, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)
		helper.PanicIfError(err)
		datas = append(datas, data)
	}
	return datas
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, where string) []entity.ProdukEntity {
	SQL := fmt.Sprintf("select id,nama,kategori,stok,created_by,created_time,updated_by,IFNULL(updated_time,'') as updated_time from produks where %s", where)
	fmt.Println(SQL)
	var datas []entity.ProdukEntity
	var data entity.ProdukEntity
	row, err := tx.QueryContext(ctx, SQL)
	fmt.Println(err, row)
	helper.PanicIfError(err)
	// fmt.Print(row)

	for row.Next() {
		err := row.Scan(&data.RowId, &data.Nama, &data.Kategori, &data.Stok, &data.CreatedBy, &data.CreatedTime, &data.UpdatedBy, &data.UpdatedTime)
		fmt.Println(err)

		helper.PanicIfError(err)
		// data.CreatedTime = helper.ConvertDateTime(data.CreatedTime)
		// data.UpdatedTime = helper.ConvertDateTime(data.UpdatedTime)

		datas = append(datas, data)
	}
	return datas
}
