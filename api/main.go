package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	handler2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello-2\n")
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/foo/", h)
	http.HandleFunc("/bar/", handler2)

	// サーバを起動
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func h (w http.ResponseWriter, _ *http.Request){
	io.WriteString(w, "Hello-1\n")
}