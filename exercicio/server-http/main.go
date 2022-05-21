package main

import (
	"log"
	"net/http"
)

func main() {
	fileSever := http.FileServer(http.Dir("./static")) // local no qual sera procurado o arquivo para ler
	http.Handle("/", fileSever)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", heloHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
