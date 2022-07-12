package main

import (
	"io/ioutil"
	"log"
	"net/http"
)
func sayhello(w http.ResponseWriter,r *http.Request){
	html,err:=ioutil.ReadFile("./hello.html")
	if err != nil {
		log.Println(err)
		return
	}
	_,_=w.Write(html)
}
func main(){
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Println(err)
		return
	}
}