package admin

import (
	"github.com/kataras/iris/v12"
	"log"
	"tower/library/apgs"
	"tower/library/databases"
	"tower/repositories/models/admins"
	"tower/repositories/repo/adminrepo"
	"tower/repositories/repository"
)

type Manager struct {
	Ctx  iris.Context
	Repo repository.Repository
}

func NewManager() *Manager {
	return &Manager{
		Repo: adminrepo.NewAdminUserRepository(),
	}
}
func (g *Manager) GetTest() {
	//where := make(map[string]interface{})
	//where["id"] = 1
	//users1 := g.Repo.Select(where)
	// 可以
	result := map[string]interface{}{}
	databases.GetDB().Model(&admins.AdminUser{}).First(&result)
	log.Println(result)
	//result := make(map[string]interface{}, 1)
	//
	//err := databases.GetDB().Table("wk_admin_user").First(&result, map[string]interface{}{"id": 1}).Error
	//log.Println(err)
	//
	//g.Ctx.JSON(apgs.ApiReturn(0, "123123",result))
	//users, _ := g.Repo.SelectById("select * from wk_admin_iser where id=?", 1)
	//fmt.Println(users)
	//fmt.Println(users1)
	//
	//_, _ = g.Ctx.JSON(apgs.ApiReturn(0, "123123", apgs.Map{
	//	"Users":  users,
	//	"Users1": users1,
	//}))
	return
}

func (g *Manager) Get() {
	_, _ = g.Ctx.JSON(apgs.ApiReturn(0, "123123", nil))
	return
}
