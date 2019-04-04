// Package routers provides ...
package routers

import (
	"net/http"

	. "github.com/taigacute/DailyTasks/controllers"
)

//InitRouter bind url with handler
func InitRouter() {
	http.HandleFunc("/login/", LoginFunc)
	http.HandleFunc("/signup/", SignUpFunc)
	http.Handle("/static/", http.FileServer(http.Dir("public")))
}