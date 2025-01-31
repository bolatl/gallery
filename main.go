package main

import (
	"fmt"
	"net/http"

	"github.com/boaltl/lenslocked/controllers"
	"github.com/boaltl/lenslocked/models"
	"github.com/boaltl/lenslocked/templates"
	"github.com/boaltl/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFs(templates.FS, "home.gohtml", "tailwind-css.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "contact.gohtml", "tailwind-css.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFs(templates.FS, "faq.gohtml", "tailwind-css.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}
	sessionService := models.SessionService{
		DB: db,
	}

	userC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	userC.Templates.New = views.Must(views.ParseFs(templates.FS, "signup.gohtml", "tailwind-css.gohtml"))
	userC.Templates.SignIn = views.Must(views.ParseFs(templates.FS, "signin.gohtml", "tailwind-css.gohtml"))
	r.Get("/signup", userC.New)
	r.Get("/signin", userC.SignIn)
	r.Post("/users", userC.Create)
	r.Post("/signin", userC.ProcessSignIn)
	r.Get("/users/me", userC.User)
	r.Post("/signout", userC.ProcessSignOut)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	fmt.Println("Starting now")
	// byte slice key for our middleware
	csrfKey := "0123456789aAbBcCdDeEfFgGhHiIjJkK"
	// function that just wraps our router securing CSRF
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: fix this before deploy
		csrf.Secure(false),
	)
	http.ListenAndServe(":3000", csrfMw(r))
}
