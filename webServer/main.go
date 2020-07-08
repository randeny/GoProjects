package main

import (
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + "Mr. Gayan Randeny"
	w.Write([]byte(message))
}

func sendFile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Gayan"))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/gayan", sendFile)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
