package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("==start==")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(time.Now().String())
		writer.WriteHeader(200)
		writer.Write([]byte("http-v-" + time.Now().String()))
	})
	http.ListenAndServe(":10082", nil)
}
