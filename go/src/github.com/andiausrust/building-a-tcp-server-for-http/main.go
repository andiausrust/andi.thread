package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"any code you want in this function")
}

func main(){

	var d hotdog
	//http.ListenAndServe(":8080", d)
	fmt.Println(*&d)
	print(&d)

}

func print(a &hotdog){
	fmt.Println(a)
}