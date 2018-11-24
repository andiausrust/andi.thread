package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type person struct{
	Fname string
	Lname string
	Items []string
}

func main() {

	http.HandleFunc("/mshl", mshl)
	http.ListenAndServe(":8080", nil)
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	james := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{
			"cool", "Jaguar",
		},
	}
	j, err := json.Marshal(james)
	if err != nil{
		fmt.Println("not able to marshal")
	}
	w.Write(j)
}