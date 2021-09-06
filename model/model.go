package model

import "gorm.io/gorm"

type User struct {
	ID           uint       `json:"id" gorm:"primarykey"`
	Email        string     `json:"email" gorm:"unique; not null"`
	Password     string     `json:"password" gorm:"not null"`
	Name         string     `json:"name" gorm:"not null"`
	Lastname     string     `json:"lastname" gorm:"not null"`
	Phone        []Phone    `json:"phone"`
	RoleCode     string     `json:"roleCode" gorm:"not null"`
	Role         Role       `json:"role" gorm:"references:Code"`
	DepartmentID int        `json:"departmentID" gorm:"not null"`
	Department   Department `json:"department" gorm:"foreignKey:DepartmentID"`
	Salary       float64    `json:"salary" gorm:"not null"`
}

type Phone struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	PhoneNo string `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
}

type Department struct {
	gorm.Model
	Title string `json:"title" gorm:"unique"`
}

type Role struct {
	gorm.Model
	Code  string `json:"code" gorm:"unique"`
	Title string `json:"title"`
}
