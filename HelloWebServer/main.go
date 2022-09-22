package main

import (
	"log"
	"net/http"
	"os"
)

// HelloHandler handles requesst for the "/hello" resources
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Web!\n"))
}

func main() {
	//get the value of the ADDR(address) enviroment variable
	addr := os.Getenv("ADDR")

	//if the address is blank, that it defaults to port 80(its going to listen to port 80 for requests).
	if len(addr) == 0 {
		addr = ":80"
	}

	//create a new mux(router)
	//the mux calls different functions for different resource paths
	mux := http.NewServeMux()

	//tell i to call the hello handler() function
	//when someone request the resource path '/hello'
	mux.HandleFunc("/hello", HelloHandler)

	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}
