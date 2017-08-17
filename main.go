package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r)
	bs, err := ioutil.ReadFile("/content/index.html")
	fmt.Printf("The repsponse is : %s\n", string(bs))
	if err != nil {
		fmt.Printf("Couldn't read index.html: %v", err)
		os.Exit(1)
	}

	io.WriteString(w, string(bs[:]))
}

func main() {
	http.HandleFunc("/", index)
	port := ":8080"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
