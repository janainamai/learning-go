package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/janainamai/study-api-go/configs"
	_ "github.com/janainamai/study-api-go/docs"
	"github.com/janainamai/study-api-go/internal/entity"
	"github.com/janainamai/study-api-go/internal/infra/database"
	"github.com/janainamai/study-api-go/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with authentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Janaina
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  License terms
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProductDatabase(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDatabase(db)
	userHanlder := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExpiresIn)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth)) // obtém token e injeta no context
		r.Use(jwtauth.Authenticator)           // verifica se o token é válido com base na secret configurada
		r.Post("/", productHandler.Create)
		r.Get("/{id}", productHandler.GetByID)
		r.Get("/", productHandler.GetAll)
		r.Put("/{id}", productHandler.Update)
		r.Delete("/{id}", productHandler.Delete)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHanlder.Create)
		r.Post("/generate_token", userHanlder.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	fmt.Printf("Listening on port 8000")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
