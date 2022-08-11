package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"perpustakaan/controller"
	"time"
)

// feature that prevent others to see file directory
type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

type neuteredReaddirFile struct {
	http.File
}

func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

//go:embed assets
var resources embed.FS

func main() {
	//directory
	dir, _ := fs.Sub(resources, "assets")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()

	//Static Route
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	images := justFilesFilesystem{http.Dir("./images/")}
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(images)))

	//Route
	mux.HandleFunc("/", controller.HandlerIndex)

	//Auth Route
	mux.HandleFunc("/signup", controller.SignupHandler)
	mux.HandleFunc("/login", controller.LoginAuth)
	mux.HandleFunc("/logout", controller.LogoutAuth)

	//Admin
	mux.HandleFunc("/admin/data_user", controller.AdminDataUserHandler)
	mux.HandleFunc("/admin/data_user/delete", controller.AdminDeleteUser)
	mux.HandleFunc("/admin/data_user/add_user", controller.AdminAddUser)
	mux.HandleFunc("/admin/data_user/edit_user", controller.AdminEditUserPage)
	mux.HandleFunc("/admin/data_user/edit_user/update", controller.AdminUpdateUser)

	//Data Buku
	mux.HandleFunc("/admin/data_buku", controller.AdminDataBuku)
	mux.HandleFunc("/admin/data_buku/delete", controller.AdminDeleteBuku)
	mux.HandleFunc("/admin/data_buku/add_buku", controller.AdminAddDataBuku)

	//Setting Server
	address := "localhost:5000"
	fmt.Printf("Server started at %s\n", address)

	server := http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}
	err := server.ListenAndServe()

	defer server.Close()
	if err != nil {
		panic(err)
	}

}
