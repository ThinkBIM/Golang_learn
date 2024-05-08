package main

import (
	"github.com/flamego/flamego"
)

func main() {
	f := flamego.Classic()
	f.Use(flamego.Logger())
	f.Get("/", func(c flamego.Context) string {
		// c.Redirect()
		return "Hello Flamego"
	})
	f.Run()
	// http.ListenAndServe()
}
