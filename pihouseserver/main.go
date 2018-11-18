package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/Jordank321/pihouse/pihouseserver/api"
	"github.com/Jordank321/pihouse/pihouseserver/db"

	"github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mssql"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var (
	debug    = flag.Bool("debug", false, "enable debugging")
	password = flag.String("password", "", "the database password")
	port     = flag.Int("port", 1433, "the database port")
	server   = flag.String("server", "", "the database server")
	user     = flag.String("user", "", "the database user")
)

func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
	})
}

//Routes sets up the base routes for the api
func Routes() *chi.Mux {
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	router := chi.NewRouter()
	router.Use(
		cors.Handler,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	// Routing for API
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/temperature", api.TemperatureRoutes(ProvideTemperaureRepository))
		r.Mount("/api/node", api.NodeRoutes(ProvideNodeRepository))
		r.Mount("/api/humidity", api.HumidityRoutes(ProvideHumidityRepository))
	})

	// Redirect to UI
	router.Get("/", http.RedirectHandler("/ui", 301).ServeHTTP)
	// Get path of current executable
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//}
	// Serve UI from wwwroot
	//filesDir := filepath.Join(dir, "wwwroot")
	//log.Printf(filesDir)
	FileServer(router, "/ui", http.Dir("./wwwroot/"))

	return router
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	// if path != "/" && path[len(path)-1] != '/' {
	// 	r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
	// 	path += "/"
	// }
	//path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf(r.RequestURI)
		fs.ServeHTTP(w, r)
	}))
}

func main() {
	viper.AutomaticEnv()
	router := Routes()
	dbret, err := ProvideDB()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(dbret)
	log.Fatal(http.ListenAndServe(":1337", router))
}
