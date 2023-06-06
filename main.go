package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバーの起動に失敗しました。", err)
	}
}
