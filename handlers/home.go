package handlers

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/base.html", "template/home.html")
	if err != nil {
		http.Error(w, "読み込みに失敗しました。", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
	}{
		Title: "home",
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "レンダリングに失敗しました。", http.StatusInternalServerError)
		return
	}
}
