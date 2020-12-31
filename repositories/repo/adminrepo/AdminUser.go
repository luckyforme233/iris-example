package adminrepo

import (
	"tower/library/databases"
	"tower/repositories/models/admins"
	"tower/repositories/repository"
)

func NewAdminUserRepository() repository.Repository {
	newRepository, _ := repository.NewRepository(
		databases.GetDB().Model(&admins.AdminUser{}),
	)
	return newRepository
}
