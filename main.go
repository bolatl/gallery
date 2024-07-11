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

	tpl := views.Must(views.ParseFs(templates.FS, "home.gohtml", "tailwind-css.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "contact.gohtml", "tailwind-css.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind-css.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	userC := controllers.Users{}
	userC.Templates.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind-css.gohtml"))
	r.Get("/signup", userC.New)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	fmt.Println("Starting now")
	http.ListenAndServe(":3000", r)
}
