package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/exception"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/helper"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/repository"
)

type UserService interface {
	VerifyCredential(ctx context.Context, username string, password string) (interface{}, string)
	Insert(ctx context.Context, user entity.UserEntity, username string, role string) string
	Register(ctx context.Context, user entity.UserEntity) string
	Update(ctx context.Context, user entity.UserEntity) string
	Delete(ctx context.Context, username string, role string) string
	FindByUsername(ctx context.Context, username string) interface{}
	FindAll(ctx context.Context, page int, perpage int, filter string, order string) []entity.UserEntity
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
func (service *UserServiceImpl) VerifyCredential(ctx context.Context, username string, password string) (interface{}, string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	res := service.UserRepository.VerifyCredential(ctx, tx, username, password)
	if len(res) == 1 {
		return true, res[0].Role
	}

	return false, ""
}
func (service *UserServiceImpl) Insert(ctx context.Context, user entity.UserEntity, username string, role string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	if role != "admin" {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "role anda bukan admin"
		message.DescGlob = "role anda bukan admin"
		message.FieldName = "role"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	}

	insertData := service.UserRepository.Insert(ctx, tx, user, username)
	if insertData == "gagal" {
		getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("username='%s'", user.Username))
		if len(getDetail) > 0 {
			var objectMessage []exception.BadRequestError
			var message exception.BadRequestError
			message.Desc = "username sudah ada"
			message.DescGlob = "username sudah ada"
			message.FieldName = "username"
			objectMessage = append(objectMessage, message)
			panic(exception.NewBadRequestError(objectMessage))

		}
	}
	return insertData

}

func (service *UserServiceImpl) Register(ctx context.Context, user entity.UserEntity) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	user.Role = "user"
	getDetail := service.UserRepository.FindAll(ctx, tx, fmt.Sprintf("username='%s' FOR UPDATE", user.Username))
	fmt.Println("hay")
	if len(getDetail) > 0 {
		var objectMessage []exception.BadRequestError
		var message exception.BadRequestError
		message.Desc = "username sudah ada"
		message.DescGlob = "username sudah ada"
		message.FieldName = "username"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))

	}
	insertData := service.UserRepository.Insert(ctx, tx, user, user.Username)

	return insertData

}
func (service *UserServiceImpl) Update(ctx context.Context, user entity.UserEntity) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	fmt.Println(user.Username, user.Password, user.Role, "hayasasasa")
	updateData := service.UserRepository.Update(ctx, tx, user)

	return updateData
}
func (service *UserServiceImpl) Delete(ctx context.Context, username string, role string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	var objectMessage []exception.BadRequestError
	var message exception.BadRequestError
	if role != "admin" {
		message.Desc = "role anda bukan admin"
		message.DescGlob = "role anda bukan admin"
		message.FieldName = "role"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	} else if username == "admin" {
		message.Desc = "tidak boleh di hapus"
		message.DescGlob = "tidak boleh di hapus"
		message.FieldName = "username"
		objectMessage = append(objectMessage, message)
		panic(exception.NewBadRequestError(objectMessage))
	}
	deleteData := service.UserRepository.Delete(ctx, tx, username)
	return deleteData
}
func (service *UserServiceImpl) FindByUsername(ctx context.Context, username string) interface{} {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	fmt.Println(username)
	getData := service.UserRepository.FindByUsername(ctx, tx, username)
	if len(getData) > 0 {
		return getData[0]
	} else {
		return nil
	}
}
func (service *UserServiceImpl) FindAll(ctx context.Context, page int, perpage int, filter string, order string) []entity.UserEntity {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)
	if filter == "" {
		filter = "1=1"
	}
	if order == "" {
		order = "username asc"
	}
	if page > 0 {
		page = page - 1
	}
	where := fmt.Sprintf("%s order by %s LIMIT %d,%d", filter, order, page, perpage)
	getData := service.UserRepository.FindAll(ctx, tx, where)
	return getData
}
