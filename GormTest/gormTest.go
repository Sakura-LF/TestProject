package GormTest

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"log/slog"
	"os"
	"time"
)

var DB *gorm.DB

//type User struct {
//	gorm.Model
//	Name     string
//	Age      int
//	Birthday time.Time
//	Active   bool
//}
//type Order struct {
//	gorm.Model
//	UserID  int
//	GoodsID int
//}

func DBConnection() {
	// 终端打印输入 sql 执行记录
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             300 * time.Millisecond, // 慢查询 SQL 阈值
			Colorful:                  true,                   // 是否启动彩色打印
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Error, // Log lever
		},
	)
	db, err := gorm.Open(mysql.Open("root:mysql_TcKPjC@tcp(192.168.233.128:3306)/testDB?parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
		//DisableForeignKeyConstraintWhenMigrating: true,
	})
	DB = db
	if err != nil {
		panic("failed to connect database")
	} else {
		slog.Info("mysql connect success")
	}
	err = db.AutoMigrate(&User{}, &Order{})
	if err != nil {
		slog.Error("mysql autoMigrate failed")
	}
	fmt.Println("mysql autoMigrate success")
}
