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

	user := entity.User{}
	var resp []entity.User
	for res.Next() {
		err := res.Scan(&user.ID, &user.Email, &user.Password, &user.Nama, &user.Role, &user.JK, &user.NoTelp, &user.Alamat)
		helper.PanicIfError(err)

		resp = append(resp, user)
	}

	data.UserData = resp

	fmt.Println(data.UserData)

	defer con.Close()

	if data.SessionData.Auth != true && data.SessionData.Role != "admin" {
		http.Redirect(w, r, "/", 303)
	}

	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/admin/user_data.gohtml"))
	err = t.ExecuteTemplate(w, "base.gohtml", &data)
	helper.PanicIfError(err)

}
