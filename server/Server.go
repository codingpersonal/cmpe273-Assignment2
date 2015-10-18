package main

import (
    "net/http"
    "log"
)

func main() {
	router := NewRouter()
    log.Fatal(http.ListenAndServe(":8081", router))
}

