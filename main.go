package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var htmlStr string

// go run main.go を実行してmain()から処理が始まる
func main() {
	fmt.Println("start")

	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}
	htmlStr = string(data)

	http.HandleFunc("/", showScreen)
	http.HandleFunc("/add_memo", addMemo)
	http.ListenAndServe(":8080", nil)
}

func showScreen(w http.ResponseWriter, r *http.Request) {
	/// fmt.Fprintln(w, "<html><h1>Hello</h1></html>")
	fmt.Fprintln(w, htmlStr)
}

type Memo struct {
	ID        string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var memos map[string]Memo

// url -X POST -H "Content-Type: application/json" -d '{"ID":"tio"}' localhost:8080/add_memo
func addMemo(w http.ResponseWriter, r *http.Request) {
	m := &Memo{}
	if err := json.NewDecoder(r.Body).Decode(m); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, m.ID)
}
