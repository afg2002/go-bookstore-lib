package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"
	"strings"

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
		err = result.Scan(&book.ID, &book.Cover, &book.Judul, &book.Harga, &book.Pengarang, &book.Kategori, &book.Penerbit, &book.Tahun, &book.Stok, &book.Deskripsi)
		helper.PanicIfError(err)

		resp = append(resp, book)
	}
	data.BookData = resp

	data.SessionData.Title = "Go Perpus"
	data.SessionData.ID = session.Values["id"]
	data.SessionData.Auth = session.Values["auth"]
	data.SessionData.Email = session.Values["email"]
	data.SessionData.Name = session.Values["name"]
	data.SessionData.Role = session.Values["role"]
	data.SessionData.Message = session.Values["message"]
	err = t.ExecuteTemplate(w, base, &data)
	helper.PanicIfError(err)

}

func UserCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		idUser := strings.TrimPrefix(r.URL.Path, "/cart/")

		con := db.ConnectionDB()

		sql := "SELECT cart.id_cart,buku.judul,buku.cover_buku,cart.total,buku.harga FROM cart INNER JOIN buku ON cart.id_buku = buku.id_buku WHERE id_user = ?;"

		res, err := con.Query(sql, idUser)
		helper.PanicIfError(err)
		cart := entity.Cart{}
		var cartResp []entity.Cart
		for res.Next() {
			err := res.Scan(&cart.IdCart, &cart.JudulBuku, &cart.CoverBuku, &cart.TotalPerItem, &cart.Harga)
			helper.PanicIfError(err)

			cartResp = append(cartResp, cart)
		}
		defer con.Close()

		json.NewEncoder(w).Encode(cartResp)

	} else if r.Method == http.MethodDelete {

		idCart := strings.TrimPrefix(r.URL.Path, "/cart/")

		con := db.ConnectionDB()
		sql := "DELETE FROM cart WHERE id_cart = ?"

		_, err := con.Exec(sql, idCart)
		helper.PanicIfError(err)
	} else {
		//Jika method selain post maka tidak diperbolehkan
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Books(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		id := strings.TrimPrefix(r.URL.Path, "/books/")

		// Buka koneksi dan Query Select
		con := db.ConnectionDB()
		sql := "SELECT * FROM buku WHERE id_buku = ?"

		res, err := con.Query(sql, id)

		helper.PanicIfError(err)
		defer res.Close()

		book := entity.Book{}
		var resp []entity.Book
		for res.Next() {
			err := res.Scan(&book.ID, &book.Cover, &book.Judul, &book.Harga, &book.Pengarang, &book.Kategori, &book.Penerbit, &book.Stok, &book.Tahun, &book.Deskripsi)
			helper.PanicIfError(err)

			resp = append(resp, book)
		}

		defer con.Close()

		json.NewEncoder(w).Encode(resp)
	} else if r.Method == http.MethodPost {

		id := strings.TrimPrefix(r.URL.Path, "/books/")
		splitId := strings.Split(id, "/")
		log.Print(splitId)
		con := db.ConnectionDB()
		sql := "SELECT * FROM cart WHERE id_buku = ? AND id_user = ?"

		res, err := con.Query(sql, splitId[0], splitId[1])

		helper.PanicIfError(err)
		defer res.Close()

		if res.Next() {
			sql := "UPDATE cart SET total = total+1 WHERE id_buku = ? AND id_user = ?"
			_, err2 := con.Exec(sql, splitId[0], splitId[1])
			helper.PanicIfError(err2)

		} else {
			sql := "INSERT INTO cart(id_buku,id_user,total) VALUES (?,?,?)"
			_, err := con.Exec(sql, splitId[0], splitId[1], 1)
			helper.PanicIfError(err)
		}

		defer con.Close()

	}
}
