package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Option struct {
	Key   string `json:"key" gorm:"primaryKey;type:string"`
	Value string `json:"value" gorm:"type:string;"`
}
