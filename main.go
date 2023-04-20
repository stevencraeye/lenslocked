package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site - Stevenn</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:stevencraeye@gmail.com\">stevencraeye@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>Is there a free version? Yes! We offer a free trial.</li>
		<li>What are your support hours? We are open 24/7.</li>
	</ul>
	`)
}


func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {http.Error(w, "Page not found", http.StatusNotFound)})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}