package datasource

import (
	"github.com/micrease/micrease-core/config"
	"github.com/micrease/micrease-core/database"
	"gorm.io/gorm"
)

//管理多库连接,读写分离或多库,微服务中不建议连接多个库，如果需要的话可以在这里扩展
type DBManager struct {
	master *gorm.DB //主库
	slave  *gorm.DB //从库，读库
}

var dbManager DBManager

func InitDatabase() {
	config := config.GetCommonConfig()
	dbManager.master = database.Connect(config.Database)
}

func GetDB() *gorm.DB {
	return dbManager.master
}

func GetDBSlave() *gorm.DB {
	return dbManager.slave
}
