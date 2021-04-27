package admin

import (
	"github.com/kataras/iris/v12"
	"tower/app/repositories/repo/adminrepo"
	"tower/library/apgs"
)

type Manager struct {
	Ctx  iris.Context
	Repo *adminrepo.AdminUserRepository
}

func NewManager() *Manager {
	return &Manager{
		Repo: adminrepo.NewAdminUserRepository(),
	}
}
func (g *Manager) getTest() {
	where := make(map[string]interface{})
	where["id"] = 1
	users1 := g.Repo.Select(where)
	g.Ctx.JSON(apgs.ApiReturn(0, "123123", users1))
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
