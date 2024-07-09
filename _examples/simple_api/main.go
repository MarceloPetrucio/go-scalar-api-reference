package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
)

// @title           Simple API
// @version         1.0
// @description     Example use of scalar beautiful api
// @termsOfService  http://swagger.io/terms/

// @contact.name   Marcelo Petrucio
// contact.url    https://marcelopetrucio.dev
// @contact.email  marcelo.petrucio43@gmail.com

// @BasePath  /
func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", create)

	router.Get("/reference", func(w http.ResponseWriter, r *http.Request) {
		t := true
		f := false

		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode:   &t,
			IsEditable: &f,
		})
		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	})

	fmt.Printf("Starting web server on port :8000")
	http.ListenAndServe(":8000", router)
}
