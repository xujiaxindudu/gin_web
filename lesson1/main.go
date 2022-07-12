package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter,r *http.Request){
	_,err:=fmt.Fprintln(w,"dudu")
	if err != nil {
		log.Println(err)
		return
	}
}
func main(){
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Println(err)
		return
	}
}