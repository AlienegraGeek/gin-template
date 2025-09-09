package db

import (
	"fmt"
	"gin-template/conf"
	"gin-template/global"
	"gin-template/logger"
	"gin-template/model"
	"gin-template/utils"
	"github.com/jinzhu/configor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

func Initialize() {
	projectPath := utils.GetProjectPath()
	err := configor.Load(&conf.Config, projectPath+"/conf.yaml")
	if err != nil {
		logger.Log.Fatal("加载配置文件失败: ", err)
		return
	}
	logger.Log.Info("配置加载成功:", conf.Config)
}

// ConnectDB connect to db
func ConnectDB() {
	allModels := []interface{}{
		model.User{}}

	dbConfig := conf.Config.DB
	if dbConfig.Port == "" {
		logger.Log.Fatal("数据库端口未配置")
	}
	port, err := strconv.ParseUint(dbConfig.Port, 10, 32)
	if err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"PortValue": dbConfig.Port,
		}).Fatal("解析数据库端口失败")
	}

	sqlLog := dblogger.New(log.New(os.Stdout, "[SQL] ", log.LstdFlags), dblogger.Config{
		SlowThreshold: 200 * time.Millisecond,
		//LogLevel:                  logger.Info,
		LogLevel:                  dblogger.Error,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", dbConfig.Host, port, dbConfig.User, dbConfig.Password, dbConfig.Name)
	if global.DB, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true, // 开启自动更新UpdatedAt字段
			Logger:                                   sqlLog,
		}); err != nil {
		logger.Log.WithFields(map[string]interface{}{
			"DSN": fmt.Sprintf("host=%s port=%d dbname=%s user=%s", dbConfig.Host, port, dbConfig.Name, dbConfig.User),
		}).Fatal("连接数据库失败: " + err.Error())
	}

	// 创建表
	for _, m := range allModels {
		if !global.DB.Migrator().HasTable(m) {
			if err = global.DB.AutoMigrate(m); err != nil {
				logger.Log.WithFields(map[string]interface{}{
					"Model": fmt.Sprintf("%T", m),
				}).Fatal("数据库迁移失败: " + err.Error())
			}
		}
	}

	//设置时区
	if err := global.DB.Exec("SET TIME ZONE 'Asia/Shanghai'").Error; err != nil {
		logger.Log.Fatal("设置数据库时区失败: ", err)
	}

	logger.Log.Info("Database Migrated Successfully")
}
