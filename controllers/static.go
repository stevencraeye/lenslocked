package controllers

import (
	"net/http"

	"github.com/stevencraeye/lenslocked/views"
)


func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct{
		Question string
		Answer string 
	}{
		{
			Question: "Is there a free version?",
			Answer: "Yes! We offer free trial for 30 days.",
		},
		{
			Question: "What are your support hours?",
			Answer: "24/7",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}