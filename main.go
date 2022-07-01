package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"perpustakaan/controller"
	"time"
)

//go:embed assets
var resources embed.FS

func main() {
	//directory
	dir, _ := fs.Sub(resources, "assets")
	fileServer := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()

	//Route
	mux.HandleFunc("/", controller.HandlerIndex)

	//Auth Route
	mux.HandleFunc("/signup", controller.SignupHandler)
	mux.HandleFunc("/login", controller.LoginAuth)
	mux.HandleFunc("/logout", controller.LogoutAuth)

	//Static Route
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Admin
	mux.HandleFunc("/admin/data_user", controller.AdminDataUserHandler)

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
