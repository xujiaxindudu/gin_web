package main

import (
	"html/template"
	"log"
	"net/http"
)


func demo1(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err != nil {
		log.Println("parse2 failed,err:",err)
		return
	}
	name:="dudu"
	err=t.Execute(w,name)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/tmplDemo",demo1)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}