package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username" gorm:"unique;type:string"`
	Password    string `json:"password" gorm:"not null;type:string;"`
	DisplayName string `json:"displayName" gorm:"type:string;"`
	Role        string `json:"role" gorm:"type:string;default:common"`    // admin, common
	Status      string `json:"status" gorm:"type:string;default:enabled"` // enabled, disabled
}

func (user *User) Insert() error {
	var err error
	err = DB.Create(user).Error
	return err
}

func (user *User) Update() error {
	var err error
	err = DB.Model(user).Updates(user).Error
	return err
}

func (user *User) Delete() error {
	var err error
	err = DB.Delete(user).Error
	return err
}

func (user *User) ValidateAndFill() {
	// When querying with struct, GORM will only query with non-zero fields,
	// that means if your field’s value is 0, '', false or other zero values,
	// it won’t be used to build query conditions
	DB.Where(&user).First(&user)
}
