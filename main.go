package main

import (
	"flag"
	"log"
	"net/http"
	"pihouse/api/temperature"

	"github.com/go-chi/chi/middleware"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "", "the database password")
	port     = flag.Int("port", 1433, "the database port")
	server   = flag.String("server", "", "the database server")
	user     = flag.String("user", "", "the database user")
)

//Routes sets up the base routes for the api
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/temperature", temperature.Routes())
	})

	return router
}

func main() {
	router := Routes()

	// walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// 	log.Panicf("%s %s\n", method, route)
	// 	return nil
	// }

	// if err := chi.Walk(router, walkFunc); err != nil {
	// 	log.Panicf("Logging err: %s\n", err.Error())
	// }

	log.Fatal(http.ListenAndServe(":8080", router))
}
