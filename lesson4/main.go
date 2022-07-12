package main

import (
	"html/template"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter,r *http.Request){
	//1.定义模板
	//2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Println("Parse template failed,err:",err)
		return
	}
	//3.渲染模板
	name:="dudu"
	err=t.Execute(w,name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func main() {
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
