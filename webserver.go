package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	fmt.Println("test")
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Welcome to my website!")
	//})
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
	time.Sleep(50 * time.Second)
}
