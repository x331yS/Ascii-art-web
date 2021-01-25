package main

import (
	"./pkg"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/favicon.ico", pkg.FaviconHandler)
	http.HandleFunc("/", pkg.Handler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	fmt.Println("Listening at localhost:6969\nHttp Status :", http.StatusOK)
	pkg.Openbrowser("http://localhost:6969")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		log.Fatal(err)
	}
}
