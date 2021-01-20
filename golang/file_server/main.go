package main

import (
	"log"
	"net/http"
)

func main() {
	dirRoot := "/usr/local/"
	var orig = http.StripPrefix("/", http.FileServer(http.Dir(dirRoot)))
	var wrapped = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		orig.ServeHTTP(w, r)
	})

	http.Handle("/", wrapped)

	//To create self signed-ssl certificate, run the following command on the command line.
	//openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./server.key -out ./server.crt
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
