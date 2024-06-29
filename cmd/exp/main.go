package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Bolat Nazar",
		Age:  20,
		Bio:  `<script>alert("Haha, you have been hx304-3!");</script>`,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}

}
