package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to Ibeji</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:support@ibeji.org\">support@ibeji.org</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We could not find the page you were looking for: please send us an email.")
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func main() {
	router := httprouter.New()
	router.GET("/hello/:name", Hello)

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", router)
}
