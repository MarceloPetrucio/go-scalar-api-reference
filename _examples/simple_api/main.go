package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// @title           Simple API
// @version         1.0
// @description     Exemple use of scalar beautfull api
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

	// router.Get("/reference", func(w http.ResponseWriter, r *http.Request) {
	// 	htmlContent, err := scalar.ApiReferenceHTML(&scalar_api_reference.Options{
	// 		SpecURL: "../../docs/swagger.json",
	// 		CustomOptions: scalar_api_reference.CustomOptions{
	// 			PageTitle: "CGC API",
	// 		},
	// 		DarkMode: true,
	// 	})

	// 	if err != nil {
	// 		fmt.Printf("%v", err)
	// 	}

	// 	fmt.Fprintln(w, htmlContent)
	// })

	http.ListenAndServe(":8000", router)
}
