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

type Book struct {
	ID        int
	Cover     string
	Judul     string
	Pengarang string
	Kategori  string
	Penerbit  string
	Tahun     string
	Stok      string
}

type Data struct {
	SessionData Session
	UserData    []User
	BookData    []Book
}

type Session struct {
	Title, Auth, Email, Name, Role, Message any
}
