package migrate

import (
	"github.com/sirupsen/logrus"
	"tower/library/databases"
	"tower/repositories/models/admins"
)

var Models = []interface{}{
	&admins.Admin{},
	&admins.Permissions{},
	&admins.Roles{},
}

func AutoMigrate() {
	db := databases.GetDB()
	// 自动创建数据库
	if migerr := db.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(Models...).Error; nil != migerr {
		logrus.Fatal("auto migrate tables failed: ")
	}
}
