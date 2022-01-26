package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	args := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"root",
		"localhost",
		"3306",
		"server_user",
	)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: args,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Errorf("failed to connect database, err: %v\n", err.Error())
		return nil, err
	}

	sqlDb, _ := DB.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return DB, nil
}
