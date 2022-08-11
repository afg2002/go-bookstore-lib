package controller

import (
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

	// Get the book data
	con := db.ConnectionDB()
	sql := "SELECT * FROM buku"
	result, err := con.Query(sql)
	helper.PanicIfError(err)

	book := entity.Book{}
	var resp []entity.Book
	for result.Next() {
		err = result.Scan(&book.ID, &book.Cover, &book.Judul, &book.Pengarang, &book.Kategori, &book.Penerbit, &book.Tahun, &book.Stok)
		helper.PanicIfError(err)

		resp = append(resp, book)
	}
	data.BookData = resp

	data.SessionData.Title = "Go Perpus"
	data.SessionData.Auth = session.Values["auth"]
	data.SessionData.Email = session.Values["email"]
	data.SessionData.Name = session.Values["name"]
	data.SessionData.Role = session.Values["role"]
	data.SessionData.Message = session.Values["message"]
	err = t.ExecuteTemplate(w, base, &data)
	helper.PanicIfError(err)

}
