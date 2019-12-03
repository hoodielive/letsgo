package main

import (
	"github.com/couchbase/gocb"
	"github.com/gorilla/mux"
	"net/http"
)

var bucket *gcb.Bucket
var bucketName string

func ExpandEndpoint(w http.ResponseWriter, req *http.Request) {}
func CreateEndpoint(w http.ResponseWriter, req *http.Request) {}
func RootEndpoint(w http.ResponseWriter, req *http.Request)   {}

func main() {
	route := mux.NewRouter()
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucketName = "onibode"
	bucket, _ = cluster.OpenBucket(bucketName, "")
	router.HandleFunc("/{id}", RootEndpoint).Methods("GET")
	router.HandleFunc("/expand/", Expan).Methods("GET")
	router.HandleFunc("/create", CreateEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":12345", router))
}
