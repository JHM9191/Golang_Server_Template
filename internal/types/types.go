package types

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	FullName       string
	LoginId        string
	Password       string
	RegisteredTime uint64
}
