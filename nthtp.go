package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){

	fmt.Println("------------------------ start ------------------------")
	for k, v := range r.Header {
		fmt.Println(k, ":", v)
	}
	fmt.Println("------------------------  end  ------------------------")
	fmt.Fprintf(w, "index for goloag.")
}

func main (){
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":6006", nil)
}