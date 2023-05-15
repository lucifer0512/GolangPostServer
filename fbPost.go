package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var outfp *os.File

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outfp.Write([]byte("@\n@Gais_REC:\n@body:"))
	outfp.Write(body)
	outfp.Write([]byte("\n"))

	// do something with the request body
//	fmt.Println("Received POST request:", string(body))

	// write the response
	fmt.Fprint(w, "POST request received ", len(body), " bytes")
}

func main() {
	// create a new HTTP server
	var err error

	outfp, err = os.OpenFile("test.rec", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer outfp.Close()
	server := http.NewServeMux()

	// set the handler for POST requests
	server.HandleFunc("/fbpost", handlePostRequest)

	// start the server
	log.Fatal(http.ListenAndServe(":7788", server))
}
