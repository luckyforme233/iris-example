package adminrepo

import (
	"github.com/jinzhu/gorm"
	"tower/library/apgs"
	"tower/library/databases"
	"tower/repositories/models/adminmods"
)

type AdminUserRepository interface {
	Select(query map[string]interface{}) apgs.Map
	SelectById(query string, id int64) (adminmods.AdminUser, error)
	SelectByName(query string, name string) adminmods.AdminUser
}

type AdminUserRepo struct {
	DB *gorm.DB
}

func NewAdminUserRepository() AdminUserRepository {
	return &AdminUserRepo{
		DB: databases.GetDB(),
	}
}

func (m *AdminUserRepo) Select(query map[string]interface{}) apgs.Map {
	var result []adminmods.AdminUser
	err := m.DB.Model(&adminmods.AdminUser{}).Where(query).Scan(&result).Error
	if err != nil {
		return nil
	}
	toMap, err := apgs.ToMap(result, "json")
	if err != nil {
		return nil
	}
	return toMap
}

func (m *AdminUserRepo) SelectById(query string, id int64) (adminmods.AdminUser, error) {
	result := adminmods.AdminUser{}
	err := m.DB.Raw(query, id).Scan(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *AdminUserRepo) SelectByName(query string, name string) adminmods.AdminUser {
	result := adminmods.AdminUser{}
	m.DB.Raw(query, name).Scan(&result)
	return result
}
