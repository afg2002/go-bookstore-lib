package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"perpustakaan/controller"
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
	mux.HandleFunc("/login", controller.LoginAuth)
	mux.HandleFunc("/logout", controller.LogoutAuth)

	//Static Route
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Setting Server
	address := "localhost:5001"
	fmt.Printf("Server started at %s\n", address)

	server := http.Server{
		Addr:    address,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
