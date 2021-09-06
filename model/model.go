package model

import "gorm.io/gorm"

type User struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `gorm:"not null"`
	Lastname     string `gorm:"not null"`
	Phone        []Phone
	RoleCode     string     `gorm:"not null"`
	Role         Role       `gorm:"references:Code"`
	DepartmentID int        `gorm:"not null"`
	Salary       float64    `gorm:"not null"`
	Department   Department `gorm:"foreignKey:DepartmentID"`
}

type Phone struct {
	gorm.Model
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
