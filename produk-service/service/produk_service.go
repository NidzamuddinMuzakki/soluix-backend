package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/exception"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/helper"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/repository"
)

type UserService interface {
	Insert(ctx context.Context, user entity.ProdukEntity) string
	Update(ctx context.Context, user entity.ProdukEntity) string
	Delete(ctx context.Context, rowId int) string
	FindById(ctx context.Context, username int) interface{}
	FindAll(ctx context.Context, page int, perpage int, filter string, order string) []entity.ProdukEntity
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepo repository.UserRepository, DB *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
		DB:             DB,
	}
}

func (service *UserServiceImpl) Insert(ctx context.Context, user entity.ProdukEntity) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("nama='%s' and kategori='%s' FOR UPDATE", user.Nama, user.Kategori))

	if len(getDetail) > 0 {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "produk sudah ada"
		message.DescGlob = "produk sudah ada"
		message.FieldName = "nama"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))

	}
	insertData := service.UserRepository.Insert(ctx, tx, user)

	return insertData

}

func (service *UserServiceImpl) Update(ctx context.Context, user entity.ProdukEntity) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	fmt.Println(user.Username, user.Role, "hayasasasa")
	getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("nama='%s' and kategori='%s' and id!=%d FOR UPDATE", user.Nama, user.Kategori, user.RowId))

	if len(getDetail) > 0 {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "produk sudah ada"
		message.DescGlob = "produk sudah ada"
		message.FieldName = "nama"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))

	}
	updateData := service.UserRepository.Update(ctx, tx, user)

	return updateData
}
func (service *UserServiceImpl) Delete(ctx context.Context, rowId int) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	deleteData := service.UserRepository.Delete(ctx, tx, rowId)
	return deleteData
}
func (service *UserServiceImpl) FindById(ctx context.Context, username int) interface{} {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	fmt.Println(username)
	getData := service.UserRepository.FindById(ctx, tx, username)
	if len(getData) > 0 {
		return getData[0]
	} else {
		return nil
	}
}
func (service *UserServiceImpl) FindAll(ctx context.Context, page int, perpage int, filter string, order string) []entity.ProdukEntity {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	if filter == "" {
		filter = "1=1"
	}
	if order == "" {
		order = "nama asc"
	}
	if page > 0 {
		page = page - 1
	}
	where := fmt.Sprintf("%s order by %s LIMIT %d,%d", filter, order, page, perpage)
	getData := service.UserRepository.FindAll(ctx, tx, where)
	return getData
}
