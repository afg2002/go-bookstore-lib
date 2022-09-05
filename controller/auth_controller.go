package controller

import (
	"fmt"
	"net/http"
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"

	"golang.org/x/crypto/bcrypt"
)

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	//Cek apakah methodnya sudah post
	if r.Method == http.MethodPost {

		//Koneksi dan query database
		con := db.ConnectionDB()
		sql := "SELECT id_user,nama,role,email,password FROM user WHERE email = ?"

		//Ambil data dari form
		email := r.FormValue("email")
		password := r.FormValue("password")

		//Validasi
		rows, err := con.Query(sql, email)
		helper.PanicIfError(err)
		user := entity.User{}

		//Session
		session, errSession := store.Get(r, "session_login")
		helper.PanicIfError(errSession)

		if rows.Next() {

			err := rows.Scan(&user.ID, &user.Nama, &user.Role, &user.Email, &user.Password)
			helper.PanicIfError(err)
			comparePassword := helper.ComparePassword(user.Password, password)

			if comparePassword {
				fmt.Println("Masuk ")
				session.Values["auth"] = true
				session.Values["id"] = user.ID
				session.Values["email"] = user.Email
				session.Values["role"] = user.Role
				session.Values["name"] = user.Nama
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				session.Values["message"] = "Email Atau Password Anda Salah."
				session.Options.MaxAge = 3
				session.Save(r, w)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		} else {
			session.Values["message"] = "Email Belum Terdaftar."
			session.Options.MaxAge = 3
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		//Jika method selain post maka tidak diperbolehkan
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func LogoutAuth(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_login")
	session.Options.MaxAge = -1
	session.Save(r, w)
	//Redirect
	http.Redirect(w, r, "/", http.StatusFound)
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		session, _ := store.Get(r, "session_login")
		con := db.ConnectionDB()
		query_select := "SELECT email FROM user WHERE email = ?"
		user := entity.User{
			Email: r.FormValue("email"),
		}
		rows, err := con.Query(query_select, user.Email)
		helper.PanicIfError(err)
		if rows.Next() {
			session.Values["message"] = "Email sudah terpakai."
			session.Options.MaxAge = 1
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			query := "INSERT INTO user(email,password,nama,role,jenis_kelamin,no_telp,alamat) VALUES (?,?,?,?,?,?,?)"
			password := r.FormValue("password")
			hashedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			helper.PanicIfError(err2)
			user := entity.User{
				Email:    r.FormValue("email"),
				Password: string(hashedPassword),
				Nama:     r.FormValue("name"),
				Role:     "anggota",
				JK:       r.FormValue("gender"),
				NoTelp:   "+62" + r.FormValue("no_telp"),
				Alamat:   r.FormValue("address"),
			}
			result, err := con.Exec(query, user.Email, user.Password, user.Nama, user.Role, user.JK, user.NoTelp, user.Alamat)
			helper.PanicIfError(err)
			id, err := result.LastInsertId()
			helper.PanicIfError(err)
			user.ID = int(id)

			//Redirect
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		//Jika method selain post maka tidak diperbolehkan
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
