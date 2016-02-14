package main

import (
	"net/http"
	"./my_app"
)

func main() {
	http.ListenAndServe(":3000", my_app.NewMux())
}
