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
	// Accept-Language:[en-US,en;q=0.9] Cache-Control:[max-age=0] Connection:[keep-alive] Cookie:[session=eyJfZnJlc2giOmZhbHNlLCJjc3JmX3Rva2VuIjoiNDNhM2JjZGM1ZmNiMjY3ZDkxN2ExODY0YzE5ZTBiMTZkMzg3YTBkMSJ9.Xy3cww.Y2VJU-Kafdj6BC0ffQYB4jiGavI]
	// Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1]
	// User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36]]
	// Body:{} GetBody:<nil> ContentLength:0 TransferEncoding:[] Close:false
	// Host:localhost:8080 Form:map[] PostForm:map[] MultipartForm:<nil> Trailer:map[] RemoteAddr:[::1]:56923
	// RequestURI:/apples TLS:<nil> Cancel:<nil> Response:<nil> ctx:0xc000146380}

	fmt.Printf("%+v\n", r.URL) // /apples \n /favicon.ico

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	// going to http://localhost:8080/apples shows: Hi there, I love apples!
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
