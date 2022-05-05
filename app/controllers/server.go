package controllers

import (
	"net/http"

	"github.com/bigmuramura/simple-todo/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)

	return http.ListenAndServe(":"+config.Config.Port, nil) // 第2引数の nil はマルチプレクサの指定で通常はデフォルトのマルチプレクサを使うため nil を指定する。登録されていないURLにアクセスされたら403を返す。
}
