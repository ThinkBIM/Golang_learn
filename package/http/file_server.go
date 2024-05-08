package main

import "net/http"

func main() {
	server := http.FileServer(http.Dir("."))

	_ = http.ListenAndServe(":8080", server)
}
