package controller

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"perpustakaan/db"
	"perpustakaan/entity"
	"perpustakaan/helper"
	"strconv"
	"strings"
)

func AdminDeleteUser(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()

	idUserParam := r.URL.Query().Get("id")
	sql := "DELETE FROM user WHERE id_user =  ?"
	_, err := con.Exec(sql, idUserParam)
	helper.PanicIfError(err)

	//Redirect
	http.Redirect(w, r, "/admin/data_user", 303)
}

func AdminAddUser(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "/admin/data_user", 303)
}

func AdminEditUserPage(w http.ResponseWriter, r *http.Request) {
	con := db.ConnectionDB()

	idUserParam := r.URL.Query().Get("id")

	sql := "SELECT * FROM user WHERE id_user = ? "

	res, err := con.Query(sql, idUserParam)

	defer res.Close()
	helper.PanicIfError(err)

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
		http.Redirect(w, r, "/", 303)
	}

	t := template.Must(template.ParseFiles("./views/base.gohtml", "./views/admin/edit_user.gohtml"))
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
	http.Redirect(w, r, "/admin/data_user", 303)
}
