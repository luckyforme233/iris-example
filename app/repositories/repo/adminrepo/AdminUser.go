package adminrepo

import (
	"fmt"
	"tower/app/repositories/models/admins"
	"tower/app/repositories/repository"
	"tower/library/databases"
)

type AdminUserRepository struct {
	repository.Repository
	admins.Admin
}

func NewAdminUserRepository() *AdminUserRepository {
	var model = admins.Admin{}
	newRepository, _ := repository.NewRepository(
		databases.GetDB().Model(&model),
	)
	return &AdminUserRepository{
		newRepository,
		model,
	}
}

func (a AdminUserRepository) Test() {
	fmt.Println("Not in Base")
}
