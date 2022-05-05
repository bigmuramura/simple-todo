package controllers

import (
	"log"
	"net/http"

	"github.com/bigmuramura/simple-todo/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // Method で処理を振り分けできる
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" { // 入力フォームで登録したとき
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", http.StatusFound) // ユーザ登録成功後、TOPページへリダイレクト
	}
}
