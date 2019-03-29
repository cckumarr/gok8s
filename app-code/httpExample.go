package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",sayHello)
	http.ListenAndServe(":8080",nil)
}

func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Println("go server started.")
	w.Write([]byte("hello there"))
}
