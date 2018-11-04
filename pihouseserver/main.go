package main

import (
	"flag"
	"log"
	"net/http"
	"pihouse/pihouseserver/api"
	"pihouse/pihouseserver/db"

	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mssql"

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
		r.Mount("/api/temperature", api.TemperatureRoutes(ProvideTemperaureRepository))
		r.Mount("/api/node", api.NodeRoutes(ProvideNodeRepository))
	})

	return router
}

func main() {
	viper.AutomaticEnv()
	router := Routes()
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(dbret)
	log.Fatal(http.ListenAndServe(":8080", router))
}
