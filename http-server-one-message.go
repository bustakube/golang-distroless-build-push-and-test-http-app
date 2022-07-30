// This program runs a webserver on port 8080 and responds 
// to any request for / with a message "success".

package main

import (
	"fmt"
	"net/http"
)

func env(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w,"success")
}

func main() {

	http.HandleFunc("/", env)
	http.ListenAndServe(":8080", nil)
}

