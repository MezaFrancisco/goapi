package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PubNub Rocks!!!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ticker(version string) {
	for {
		time.Sleep(20 * time.Second)
		log.Println("Go API Version:", version)
	}

}

func main() {
	version := "1"

	log.Println("Version:", version)
	go ticker(version)
	handleRequests()
	log.Println("Hello world!")

}
