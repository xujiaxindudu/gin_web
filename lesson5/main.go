package main

import (
	"html/template"
	"log"
	"net/http"
)
type User struct {
	Name string
	Gender string
	Age int
}
func sayhello(w http.ResponseWriter,r *http.Request){
	//1.定义模板
	//2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Println("Parse failed,err:",err)
		return
	}
	//3.渲染模板
	u1:=User{
		Name: "嘟嘟",
		Gender: "女",
		Age: 3,
	}
	m1:=map[string]interface{}{
		"name":"嘟嘟",
		"Gender":"女",
		"Age":18,
	}
	hobbyList:=[]string{
		"篮球",
		"足球",
		"双色球",
	}
	err=t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
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
