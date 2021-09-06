package database

import (
	"github.com/subAlgo/userFiber/model"
	_ "gorm.io/driver/postgres"
)

func Migration() (err error) {

	if err = DB.Migrator().DropTable(&model.User{}, &model.Department{}, &model.Role{}, &model.Phone{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&model.User{}, &model.Department{}, &model.Role{}, &model.Phone{}); err != nil {
		return err
	}
	return nil
}
