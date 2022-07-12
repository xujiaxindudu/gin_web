package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter,r*http.Request){
	//定义模板(模板继承的方式)
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/index.tmpl")
	if err != nil {
		log.Println("parse failed,err:",err)
		return
	}
	//渲染模板
	name:="嘟嘟"
	err=t.ExecuteTemplate(w,"index.tmpl",name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func home(w http.ResponseWriter,r*http.Request){
	//定义模板(模板继承的方式)
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/home.tmpl")
	if err != nil {
		log.Println("parse failed,err:",err)
		return
	}
	//渲染模板
	name:="一一"
	err=t.ExecuteTemplate(w,"home.tmpl",name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}