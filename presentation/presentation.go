package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/get-key?sku=%s", string(os.Getenv("CATALOG_API_URL")), r.URL.Query().Get("sku"))
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	io.WriteString(w, string(body))
}

func main() {
	if os.Getenv("CATALOG_API_URL") == "" {
		log.Fatalln("Missing CATALOG_API_URL environment variable ")
	}
	http.HandleFunc("/", ApiHandler)
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
