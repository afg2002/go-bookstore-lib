package controller

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()

	idUserParam := r.URL.Query().Get("id")
	sql := "DELETE FROM user WHERE id_user =  ?"
	_, err := con.Exec(sql, idUserParam)
	helper.PanicIfError(err)

	//Redirect
	http.Redirect(w, r, "/admin/data_user", http.StatusSeeOther)
}

func AdminAddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		con := db.ConnectionDB()
		query := "INSERT INTO user(email,password,nama,role,jenis_kelamin,no_telp,alamat) VALUES (?,?,?,?,?,?,?)"
		password := r.FormValue("userPass")
		hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		helper.PanicIfError(err2)
		user := entity.User{
			Email:    r.FormValue("userEmail"),
			Password: string(hashedPassword),
			Nama:     r.FormValue("userNama"),
			Role:     r.FormValue("userRole"),
			JK:       r.FormValue("userJK"),
			NoTelp:   "+62" + r.FormValue("userNoTelp"),
			Alamat:   r.FormValue("userAlamat"),
		}
		result, err := con.Exec(query, user.Email, user.Password, user.Nama, user.Role, user.JK, user.NoTelp, user.Alamat)
		helper.PanicIfError(err)
		id, err := result.LastInsertId()
		helper.PanicIfError(err)
		user.ID = int(id)

		//Redirect
		http.Redirect(w, r, "/admin/data_user", http.StatusSeeOther)
	} else {
		//Jika method selain post maka tidak diperbolehkan
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func AdminEditUserPage(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()

	idUserParam := r.URL.Query().Get("id")

	sql := "SELECT * FROM user WHERE id_user = ? "

	res, err := con.Query(sql, idUserParam)

	helper.PanicIfError(err)
	defer res.Close()

	user := entity.User{}
	var resp []entity.User
	if res.Next() {
		err := res.Scan(&user.ID, &user.Email, &user.Password, &user.Nama, &user.Role, &user.JK, &user.NoTelp, &user.Alamat)
		helper.PanicIfError(err)

		resp = append(resp, user)
	}

	data.UserData = resp

	//Delete NoTelp Prefix
	delNoTelpPrefix := strings.TrimLeft(resp[0].NoTelp, "+62")
	resp[0].NoTelp = delNoTelpPrefix

	defer con.Close()

	if data.SessionData.Auth != true && data.SessionData.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/admin/user/edit_user.gohtml"))
	err = t.ExecuteTemplate(w, "base.gohtml", &data)
	helper.PanicIfError(err)
}

func AdminUpdateUser(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()
	query := "UPDATE user SET email = ? ,password = ? ,nama = ? ,role = ? ,jenis_kelamin = ? ,no_telp = ?,alamat = ? WHERE id_user = ?"
	id, _ := strconv.Atoi(r.FormValue("userID"))
	password := r.FormValue("userPass")
	hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.PanicIfError(err2)
	user := entity.User{
		ID:       id,
		Email:    r.FormValue("userEmail"),
		Password: string(hashedPassword),
		Nama:     r.FormValue("userNama"),
		Role:     r.FormValue("userRole"),
		JK:       r.FormValue("userJK"),
		NoTelp:   "+62" + r.FormValue("userNoTelp"),
		Alamat:   r.FormValue("userAlamat"),
	}
	_, err := con.Exec(query, &user.Email, &user.Password, &user.Nama, &user.Role, &user.JK, &user.NoTelp, &user.Alamat, &user.ID)
	helper.PanicIfError(err)

	//Redirect
	http.Redirect(w, r, "/admin/data_user", http.StatusSeeOther)
}

func AdminDataUserHandler(w http.ResponseWriter, r *http.Request) {
	// Buka koneksi dan Query Select
	con := db.ConnectionDB()
	sql := "SELECT * FROM user"

	res, err := con.Query(sql)

	helper.PanicIfError(err)
	defer res.Close()

	user := entity.User{}
	var resp []entity.User
	for res.Next() {
		err := res.Scan(&user.ID, &user.Email, &user.Password, &user.Nama, &user.Role, &user.JK, &user.NoTelp, &user.Alamat)
		helper.PanicIfError(err)

		resp = append(resp, user)
	}

	data.UserData = resp

	defer con.Close()

	if data.SessionData.Auth != true && data.SessionData.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/admin/user/user_data.gohtml"))
	err = t.ExecuteTemplate(w, "base.gohtml", &data)
	helper.PanicIfError(err)

}

func AdminDataBuku(w http.ResponseWriter, r *http.Request) {
	// Buka koneksi dan Query Select
	con := db.ConnectionDB()
	sql := "SELECT * FROM buku"

	res, err := con.Query(sql)

	helper.PanicIfError(err)
	defer res.Close()

	book := entity.Book{}
	var resp []entity.Book
	for res.Next() {
		err = res.Scan(&book.ID, &book.Cover, &book.Judul, &book.Harga, &book.Pengarang, &book.Kategori, &book.Penerbit, &book.Tahun, &book.Stok, &book.Deskripsi)
		helper.PanicIfError(err)

		resp = append(resp, book)
	}

	data.BookData = resp

	defer con.Close()

	if data.SessionData.Auth != true && data.SessionData.Role != "admin" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/admin/book/book_data.gohtml"))
	err = t.ExecuteTemplate(w, "base.gohtml", &data)
	helper.PanicIfError(err)
}

func AdminAddDataBuku(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()
	query := "INSERT INTO buku(cover_buku,judul,harga,pengarang,kategori,penerbit,tahun,stok,deskripsi) VALUES (?,?,?,?,?,?,?,?,?)"

	r.ParseMultipartForm(10 << 20)
	file, fileHeader, err := r.FormFile("coverBuku")

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	defer file.Close()

	dest, err := os.Create("./images/" + fileHeader.Filename)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	harga, _ := strconv.Atoi(r.FormValue("hargaBuku")) // convert 'harga' from string to Integer
	buku := entity.Book{
		Cover:     fileHeader.Filename,
		Judul:     r.FormValue("judulBuku"),
		Harga:     harga,
		Pengarang: r.FormValue("pengarangBuku"),
		Kategori:  r.FormValue("kategoriBuku"),
		Penerbit:  r.FormValue("penerbitBuku"),
		Tahun:     r.FormValue("tahunTerbit"),
		Stok:      r.FormValue("stokBuku"),
		Deskripsi: r.FormValue("deskripsi"),
	}

	result, err2 := con.Exec(query, buku.Cover, buku.Judul, buku.Harga, buku.Pengarang, buku.Kategori, buku.Penerbit, buku.Tahun, buku.Stok, buku.Deskripsi)
	helper.PanicIfError(err2)
	id, err3 := result.LastInsertId()
	buku.ID = int(id)
	helper.PanicIfError(err3)

	//Redirect
	http.Redirect(w, r, "/admin/data_buku", http.StatusSeeOther)
}

func AdminDeleteBuku(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()

	idUserParam := r.URL.Query().Get("id")

	sqlSelect := "SELECT cover_buku FROM buku WHERE id_buku = ?"
	getCoverBuku, err := con.Query(sqlSelect, idUserParam)
	helper.PanicIfError(err)

	buku := entity.Book{}
	if getCoverBuku.Next() {
		err2 := getCoverBuku.Scan(&buku.Cover)
		helper.PanicIfError(err2)
	}

	if _, err := os.Stat("./images/" + buku.Cover); errors.Is(err, os.ErrNotExist) {

	} else {
		path := "./images/" + buku.Cover
		err = os.Remove(path)
		helper.PanicIfError(err)
	}

	sqlDel := "DELETE FROM buku WHERE id_buku =  ?"
	_, err3 := con.Exec(sqlDel, idUserParam)
	helper.PanicIfError(err3)

	//Redirect
	http.Redirect(w, r, "/admin/data_buku", http.StatusSeeOther)
}
