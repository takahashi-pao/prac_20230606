package main

import (
	"log"
	"net/http"

	"github.com/takahashi-pao/prac_20230606/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("サーバーの起動に失敗しました。", err)
	}
}
