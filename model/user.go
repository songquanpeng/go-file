package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username" gorm:"unique;"`
	Password    string `json:"password" gorm:"not null;"`
	DisplayName string `json:"displayName"`
	Role        int    `json:"role" gorm:"type:int;default:1"`   // admin, common
	Status      int    `json:"status" gorm:"type:int;default:1"` // enabled, disabled
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
