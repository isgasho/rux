package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kataras/muxie"
)

// install bombardier:
// 	go get -u github.com/codesenberg/bombardier
// run serve:
// 	go run ./_benchmarks/muxie
// bench test:
// 	bombardier -c 125 -n 1000000 http://localhost:3000
func main() {
	r := muxie.NewMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		_, _ = w.Write([]byte("Welcome!\n"))
	})

	r.HandleFunc("/user/:id", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.NotFound(w, r)
			return
		}

		_, _ = w.Write([]byte(muxie.GetParam(w, "id")))
	})

	fmt.Println("Server started at localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}