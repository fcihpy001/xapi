package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Phone    string
	NickName string
	Gender   int8
}

type Goods struct {
	gorm.Model
}

type Order struct {
	gorm.Model
}

type Trade struct {
	gorm.Model
}

type Cart struct {
}

type Shop struct {
}

type Address struct {
}

type Payment struct {
}
