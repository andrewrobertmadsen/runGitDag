package models

type User struct {
	Username   string
	Password   string
	Type UserType
}

type UserType string

const (
	Admin     UserType = "Admin"
	Regular   UserType = "Regular"
)
