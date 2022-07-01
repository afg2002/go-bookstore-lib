package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var base = "base.gohtml"

var data entity.Data

func HandlerIndex(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session_login")
	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/index.gohtml"))
	data.SessionData.Title = "Go Perpus"
	data.SessionData.Auth = session.Values["auth"]
	data.SessionData.Email = session.Values["email"]
	data.SessionData.Name = session.Values["name"]
	data.SessionData.Role = session.Values["role"]
	data.SessionData.Message = session.Values["message"]
	err := t.ExecuteTemplate(w, base, &data)
	helper.PanicIfError(err)
}

func AdminDataUserHandler(w http.ResponseWriter, r *http.Request) {
	// Buka koneksi dan Query Select
	con := db.ConnectionDB()
	sql := "SELECT * FROM user"

	res, err := con.Query(sql)

	defer res.Close()
	helper.PanicIfError(err)

	//Define Slice UserData

	// Buat array penampungan struct yang isinya slice user (slice of struct)
	rs := make([]*entity.User, 0)
	for res.Next() {
		rst := new(entity.User) // Buat array kosong untuk entity user
		err := res.Scan(&rst.ID, &rst.Email, &rst.Password, &rst.Nama, &rst.Role, &rst.JK, &rst.NoTelp, &rst.Alamat)
		helper.PanicIfError(err)

		// Append ke array slice of struct pada entity user (rst)
		rs = append(rs, rst)

	}

	// lalu isi semua Slice of Struct Data dengan rs
	data.UserData = rs

	if data.SessionData.Auth != true && data.SessionData.Role != "admin" {
		http.Redirect(w, r, "/", 303)
	}

	fmt.Println(data.UserData[0].Nama)
	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/admin/user_data.gohtml"))
	err = t.ExecuteTemplate(w, "base.gohtml", &data)
	helper.PanicIfError(err)

}
