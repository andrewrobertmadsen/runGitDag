package models

type User struct {
	Id       int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Username string `gorm:"not null" form:"username" json:"username"`
	Password string
	Type     UserType `gorm:"default:Regular"`
}

type UserType string

const (
	Admin   UserType = "Admin"
	Regular UserType = "Regular"
)
