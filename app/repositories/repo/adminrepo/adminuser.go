package adminrepo

import (
	"tower/app/repositories/models/admins"
	"tower/library/apgs"
	"tower/library/databases"
	"tower/library/password"
	"tower/library/repository"
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

func (a AdminUserRepository) Login(username, pass string) *apgs.Response {

	adminData := databases.DB.Model(&a.Admin).Preload("Roles").Where("username=?", username).First(&a.Admin)
	if adminData.Error != nil {
		return apgs.ApiReturn(400, "账号或密码错误", nil)
	}
	passBool := password.Compare(a.Admin.Password, pass)
	if passBool != nil {
		return apgs.ApiReturn(400, "账号或密码错误", nil)
	}
	a.Admin.Password = ""
	return apgs.ApiReturn(0, "", a.Admin)
}
