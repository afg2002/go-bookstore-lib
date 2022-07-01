package entity

import "database/sql"

type Database struct {
	DB *sql.DB
}

type User struct {
	ID       int
	Email    string
	Password string
	Nama     string
	Role     string
	JK       string
	NoTelp   string
	Alamat   string
}

type Data struct {
	SessionData Session
	UserData    []*User
}

type Session struct {
	Title, Auth, Email, Name, Role, Message any
}
