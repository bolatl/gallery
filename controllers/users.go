package controllers

import (
	"fmt"
	"net/http"

	"github.com/boaltl/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	// here r.FormValue does ParseForm by itself
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something is wrong:(", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email, Password string
	}
	data.Email, data.Password = r.FormValue("email"), r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Authentication failed", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:  "email",
		Value: data.Email,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Successfully authenticated: %+v", user)
}

func (u Users) User(w http.ResponseWriter, r *http.Request) {
	// here email is a cookie with all of its fields
	email, err := r.Cookie("email")
	if err != nil {
		fmt.Fprint(w, "The email cookie cannot be read")
		return
	}
	fmt.Fprintf(w, "Email: %+v\n", email.Value)
	// Cookie will be printed as the part of the header here
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}
