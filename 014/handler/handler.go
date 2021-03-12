package handler

import (
	"golang-webdev/014/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "I'm learning golang web",
		"content": "I'm learning golang web from online course",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
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

	// fmt.Fprintf(w, "Product Page : %d", idNumb)
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"content": idNumb,
	// }

	// data := entity.Product{ID: 1, Name: "Ertiga", Price: 220000000, Stock: 3}
	data := []entity.Product{
		{ID: 1, Name: "Ertiga", Price: 220000000, Stock: 11},
		{ID: 2, Name: "Xpander", Price: 260000000, Stock: 8},
		{ID: 3, Name: "Pajero Sport", Price: 470000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
}

func PostGetHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("This is a GET request"))

	case "POST":
		w.Write([]byte("This is a POST request"))

	default:
		http.Error(w, "Error is happening", http.StatusBadRequest)
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))

		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening", http.StatusInternalServerError)
			return
		}

		// name := r.Form.Get("name")
		// w.Write([]byte(name))

		message := r.Form.Get("message")
		w.Write([]byte(message))

		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}
