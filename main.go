package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var htmlStr string

// go run main.go を実行してmain()から処理が始まる
func main() {
	fmt.Println("start")

	http.HandleFunc("/getindexhtml", getindexHTML)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><h1>Hello</h1></html>")
}

func getindexHTML(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	htmlStr = string(data)

	fmt.Fprintln(w, htmlStr)
}
