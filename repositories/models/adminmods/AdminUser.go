package adminmods

import (
	"github.com/jinzhu/gorm"
)

type AdminUser struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
	//Name     string //姓名
	//Email    string //邮箱
	//Mobile   string //手机
	//QQ       string
	//Gender   int    //0男 1女
	//Age      int    //年龄
	//Remark   string //备注
	//Token    string `gorm:"-"`
	//Session  string `gorm:"-"`
}
