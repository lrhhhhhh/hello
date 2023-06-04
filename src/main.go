package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        x := rand.Intn(1<<30)
        fmt.Fprintf(w, "Hello, No.%d", x)
    })

    if err := http.ListenAndServe(":8888", nil); err != nil {
        panic(err)
    }
}
