package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Golang Web Development!"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, I'm learning golang web development"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from Nintendo"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product Page : %d", idNumb)
}
