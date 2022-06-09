package controller

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"os"
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	//Cek apakah methodnya sudah post
	if r.Method == "POST" {
		//Koneksi dan query database

		con := db.ConnectionDB()
		sql := "SELECT email,password FROM user WHERE email = ?"

		//Ambil data dari form
		email := r.FormValue("email")
		password := r.FormValue("password")

		//Validasi
		rows, err := con.Query(sql, email)
		helper.PanicIfError(err)
		user := entity.User{}
		if rows.Next() {
			err := rows.Scan(&user.Email, &user.Password)
			helper.PanicIfError(err)
			comparePassword := helper.ComparePassword(user.Password, password)
			if comparePassword {
				session, _ := store.Get(r, "session_login")
				session.Values["auth"] = true
				session.Values["email"] = user.Email
				session.Save(r, w)
				http.Redirect(w, r, "/", 302)
			} else {
				http.Redirect(w, r, "/", 303)
			}
		}
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func LogoutAuth(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_login")
	session.Values["auth"] = false
	session.Save(r, w)

	//Redirect
	http.Redirect(w, r, "/", 302)
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_login")
	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/index.gohtml"))
	t.ExecuteTemplate(w, "base.gohtml", map[string]interface{}{
		"Title": "Go Perpus",
		"Auth":  session.Values["auth"],
		"Email": session.Values["email"],
	})
}
