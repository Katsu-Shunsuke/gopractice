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
	http.HandleFunc("/list_memos", listMemos)
	http.ListenAndServe(":8080", nil)
}

func showScreen(w http.ResponseWriter, r *http.Request) {
	/// fmt.Fprintln(w, "<html><h1>Hello</h1></html>")
	fmt.Fprintln(w, htmlStr)
}

// 構造体自体の定義は*をつけない
type Memo struct {
	ID        string
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Memo構造体をポインタ変数として定義している
var memos map[string]*Memo = map[string]*Memo{}

// curl -X POST -H "Content-Type: application/json" -d '{"ID":"1111","Title":"mytitle","Body":"mybody","CreatedAt":"2022-01-01T10:00:00Z","UpdatedAt":"2022-01-01T11:00:00Z"}' localhost:8080/add_memo
func addMemo(w http.ResponseWriter, r *http.Request) {
	//var m *Memo = &Memo{}
	//mがポインタ変数
	m := &Memo{}
	if err := json.NewDecoder(r.Body).Decode(m); err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	//メモをmemosに保存
	//m.IDをキーにしている
	memos[m.ID] = m

	fmt.Fprintln(w, nil)
}

// 保存してあるメモの一覧をJSONで出力する
func listMemos(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(memos)
	if err != nil {
		fmt.Fprintln(w, "error:"+err.Error())
		return
	}

	fmt.Fprintln(w, string(b))
}
