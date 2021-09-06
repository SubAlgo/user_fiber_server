package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (err error) {
	dsn := "host=localhost user=postgres password=test123456 dbname=userFiber port=9000 sslmode=disable TimeZone=Asia/Bangkok"
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}
	return nil
}
