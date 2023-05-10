package main

import (
	"fmt"
	"net/http"
)

var count int

// go run main.go を実行してmain()から処理が始まる
func main() {
	fmt.Println("start")
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", countHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Fprintln(w, count)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><h1>count</h1></html>")
}
