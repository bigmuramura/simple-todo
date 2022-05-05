package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bigmuramura/simple-todo/config"
)

// ... は過変長の値を受け取りたい
// 参考: https://qiita.com/Yashy/items/a02b74a00136dc5a42c4
// 上記ではslice同士の連結でhoge...の意味がわからず、以下のリンクを見て理解
// 参考: https://www.delftstack.com/ja/howto/go/how-to-concatenate-two-slices-in-go/
func generateHTML(w http.ResponseWriter, date interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...)) // Mustの意味を理解していない
	templates.ExecuteTemplate(w, "layout", date)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)

	return http.ListenAndServe(":"+config.Config.Port, nil) // 第2引数の nil はマルチプレクサの指定で通常はデフォルトのマルチプレクサを使うため nil を指定する。登録されていないURLにアクセスされたら403を返す。
}
