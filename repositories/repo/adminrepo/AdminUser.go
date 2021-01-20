package adminrepo

import (
	"fmt"
	"tower/library/databases"
	"tower/repositories/models/admins"
	"tower/repositories/repository"
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
