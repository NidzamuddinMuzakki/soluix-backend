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
	Insert2(ctx context.Context, q string) string

	// FindById(ctx context.Context, username int) interface{}
	FindAll(ctx context.Context, page int, perpage int, filter string, order string) ([]entity.ProdukEntity, entity.InfoList)
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

	// getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("nama='%s' and kategori='%s' FOR UPDATE", user.Nama, user.Kategori))

	// if len(getDetail) > 0 {
	// 	var objectMessage []exception.BadRequestError
	// 	var message exception.BadRequestError
	// 	message.Desc = "produk sudah ada"
	// 	message.DescGlob = "produk sudah ada"
	// 	message.FieldName = "nama"
	// 	objectMessage = append(objectMessage, message)
	// 	panic(exception.NewBadRequestError(objectMessage))

	// }

	insertData := service.UserRepository.Insert(ctx, tx, user)
	if insertData == "gagal" {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "gagal"
		message.DescGlob = "gagal"
		message.FieldName = "nama"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))

	}
	return insertData

}

func (service *UserServiceImpl) Insert2(ctx context.Context, q string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("nama='%s' and kategori='%s' FOR UPDATE", user.Nama, user.Kategori))

	// if len(getDetail) > 0 {
	// 	var objectMessage []exception.BadRequestError
	// 	var message exception.BadRequestError
	// 	message.Desc = "produk sudah ada"
	// 	message.DescGlob = "produk sudah ada"
	// 	message.FieldName = "nama"
	// 	objectMessage = append(objectMessage, message)
	// 	panic(exception.NewBadRequestError(objectMessage))

	// }

	insertData := service.UserRepository.Insert2(ctx, tx, q)
	if insertData == "gagal" {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "gagal"
		message.DescGlob = "gagal"
		message.FieldName = "nama"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))

	}
	return insertData

}

// func (service *UserServiceImpl) FindById(ctx context.Context, username int) interface{} {
// 	tx, err := service.DB.Begin()
// 	helper.PanicIfError(err)
// 	defer helper.CommitOrRollback(tx)
// 	fmt.Println(username)
// 	getData := service.UserRepository.FindById(ctx, tx, username)
// 	if len(getData) > 0 {
// 		return getData[0]
// 	} else {
// 		return nil
// 	}
// }
func (service *UserServiceImpl) FindAll(ctx context.Context, page int, perpage int, filter string, order string) ([]entity.ProdukEntity, entity.InfoList) {
	info := entity.InfoList{}
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	if filter == "" {
		filter = "1=1"
	}
	if order == "" {
		order = "id asc"
	}
	if page > 0 && perpage > 0 {
		page = (page - 1) * perpage
	} else if page > 0 {
		page = (page - 1) * 10
	}
	where := fmt.Sprintf("%s order by %s LIMIT %d,%d", filter, order, page, perpage)
	getData := service.UserRepository.FindAll(ctx, tx, where)
	if len(getData) > 0 {
		whereAll := fmt.Sprintf("%s order by %s", filter, order)
		getDataAll := service.UserRepository.FindAll(ctx, tx, whereAll)
		info.Allrec = len(getDataAll)
		info.Sentrec = len(getData)
	}
	return getData, info
}
