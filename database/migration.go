package database

import (
	"github.com/subAlgo/userFiber/model"
	"github.com/subAlgo/userFiber/pkg/password"
	_ "gorm.io/driver/postgres"
)

func Migration() (err error) {

	if err = DB.Migrator().DropTable(&model.User{}, &model.Department{}, &model.Role{}, &model.Phone{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&model.User{}, &model.Department{}, &model.Role{}, &model.Phone{}); err != nil {
		return err
	}
	// init department
	department := model.Department{Title: "Manager"}
	if err = DB.Create(&department).Error; err != nil {
		return err
	}

	roleAdmin := model.Role{Code: "1", Title: "Admin"}
	if err = DB.Create(&roleAdmin).Error; err != nil {
		return err
	}

	roleUser := model.Role{Code: "2", Title: "User"}
	if err = DB.Create(&roleUser).Error; err != nil {
		return err
	}

	pass, _ := password.Hash("123456")
	user := model.User{
		Email:        "admin@gmail.com",
		Password:     pass,
		Name:         "admin",
		Lastname:     "-",
		RoleCode:     roleAdmin.Code,
		DepartmentID: 1,
		Salary:       50000.00,
	}
	if err = DB.Create(&user).Error; err != nil {
		return err
	}

	return err
}
