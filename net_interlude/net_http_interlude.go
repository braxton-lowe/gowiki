package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r) // prints struct with field names below!!
	// &{Method:GET URL:/apples Proto:HTTP/1.1 ProtoMajor:1
	// ProtoMinor:1 Header:map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9] Accept-Encoding:[gzip, deflate, br]
	// Accept-Language:[en-US,en;q=0.9] ...
	// Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false
	// Host:localhost:8080 Form:map[] PostForm:map[]...

	fmt.Printf("%+v\n", r.URL) // /apples \n /favicon.ico

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// going to http://localhost:8080/apples shows: Hi there, I love apples!
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
