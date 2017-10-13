package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Baitoru(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	baseUrl := "http://localhost:8090/"
	response, err := http.Get(baseUrl)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%q", robots)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
