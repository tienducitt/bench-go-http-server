package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", sendMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func sendMessage(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Receive request " + time.Now().Format(time.RFC850))
	time.Sleep(3 * time.Second)
	w.WriteHeader(200)
}
