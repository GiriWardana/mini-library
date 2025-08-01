package models

type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"unique"`
    Password string
    Role     string // "admin" or "user"
}
