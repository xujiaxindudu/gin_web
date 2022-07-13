package main

import (
	"html/template"
	"log"
	"net/http"
)
func index(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.New("index.tmpl").
		Delims("{[","]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		log.Println("parse failed,err:",err)
		return
	}
	//渲染模板
	name:="dudu"
	err=t.Execute(w,name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func xss(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML{
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		log.Println("parse xss failed,err:",err)
		return
	}
	//渲染模板
	str1:="<script>alert(123)</script>"
	str2 := "<a href='https://liwenzhou.com'>liwenzhou的博客</a>"
	err=t.Execute(w,map[string]interface{}{
		"str1":str1,
		"str2":str2,

	})
	if err != nil {
		log.Println("exec xss failed,err:",err)
		return
	}

}

func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/xss",xss)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}