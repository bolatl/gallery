package main

import (
	"fmt"
	"net/http"

	"github.com/boaltl/lenslocked/controllers"
	"github.com/boaltl/lenslocked/templates"
	"github.com/boaltl/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFs(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	fmt.Println("Starting now")
	http.ListenAndServe(":3000", r)
}
