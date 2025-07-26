package controllers

import "gorm.io/gorm"

var (
	userDB   *gorm.DB
	bookDB   *gorm.DB
	borrowDB *gorm.DB
)

func SetDB(db *gorm.DB) {
	userDB = db
	bookDB = db
	borrowDB = db
}

// Or if you already have SetUserDB, SetBookDB, SetBorrowDB defined:
func SetUserDB(db *gorm.DB) {
	userDB = db
}

func SetBookDB(db *gorm.DB) {
	bookDB = db
}

func SetBorrowDB(db *gorm.DB) {
	borrowDB = db
}
