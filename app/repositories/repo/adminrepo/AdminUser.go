package adminrepo

import (
	"fmt"
	"tower/app/repositories/models/admins"
	"tower/app/repositories/repository"
	"tower/library/databases"
)

type AdminUserRepository struct {
	repository.Repository
}

func NewAdminUserRepository() *AdminUserRepository {
	newRepository, _ := repository.NewRepository(
		databases.GetDB().Model(&admins.Admin{}),
	)
	return &AdminUserRepository{
		newRepository,
	}
}

func (a AdminUserRepository) Test() {
	fmt.Println("Not in Base")
}
