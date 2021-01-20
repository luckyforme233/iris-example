package migrate

import (
	"github.com/sirupsen/logrus"
	"tower/app/repositories/models/admins"
	"tower/library/databases"
)

var Models = []interface{}{
	&admins.Admin{},
	&admins.Permissions{},
	&admins.Roles{},
}

func AutoMigrate() {
	db := databases.GetDB()
	// 自动创建数据库
	if err := db.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(Models...).Error; nil != err {
		logrus.Fatal("auto migrate tables failed: ")
	}
}
