package main

import (
	"fmt"
	"net/http"
	"student/server"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	fmt.Println("http://localhost:8080/ is listening....")
	server.Func()
	server.HandleFuncOwn()
}
