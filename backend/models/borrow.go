package models

import "time"

type Borrow struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	UserID     uint       `json:"user_id"`
	User       User       `json:"user" gorm:"foreignKey:UserID"`
	BookID     uint       `json:"book_id"`
	Book       Book       `json:"book" gorm:"foreignKey:BookID"`
	Returned   bool       `json:"returned"`
	CreatedAt  time.Time  `json:"created_at"`
	ReturnedAt *time.Time `json:"returned_at,omitempty"`
}
