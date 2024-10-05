package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, we offer trial for 7 days.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "From 9 am to 6 pm.",
		},
		{
			Question: "How can we get in touch with you?",
			Answer:   `You can email us at <a href="mailto:bolatlobakbai@gmail.com">bolatlobakbai@gmail.com</a>`,
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
