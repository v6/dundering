package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var FUCounter = 0

func sayFU(responsewriter http.ResponseWriter, request *http.Request) {
	// When someone visits the home page, (root URL, nothing after
	// port and IP address) show them an insult and add to the counter.
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println(request.Form["url_long"])
	for response_key, response_value := range request.Form {
		fmt.Println("key:", response_key)
		fmt.Println("val:", strings.Join(response_value, ""))
	}
	fmt.Fprintf(responsewriter, "OY! OY YOU! FU!") // browser
	FUCounter = FUCounter + 1
}

func assign(responsewriter http.ResponseWriter, request *http.Request) {
	// When anyone visits the computer's IP address plus /assign,
	// return the number of times someone's visited the
	// root URL (nothing after (IP address xor domain name) and port)
	// of the computer this is running on.
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println(request.Form["url_long"])
	for response_key, response_value := range request.Form {
		fmt.Println("key:", response_key)
		fmt.Println("val:", strings.Join(response_value, ""))
	}
	fmt.Fprintf(responsewriter, "Task counter has incremented to ")
	fmt.Fprintf(responsewriter, strconv.Itoa(FUCounter))
}

func main() {
	http.HandleFunc("/", sayFU)
	http.HandleFunc("/assign", assign)
	err := http.ListenAndServe(":31337", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
