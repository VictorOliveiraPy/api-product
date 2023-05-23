package main

import (
	"log"
	"net/http"

	"github.com/VictorOliveiraPy/configs"
	_ "github.com/VictorOliveiraPy/docs"
	"github.com/VictorOliveiraPy/internal/entity"
	"github.com/VictorOliveiraPy/internal/infra/database"
	"github.com/VictorOliveiraPy/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Victor Hugo
// @contact.url    https://www.linkedin.com/in/victor-hugo-3548a915a/
// @contact.email  oliveiravictordev@gmail.com

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err!= nil {
        panic(err)
    }

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err!= nil {
        panic(err)
    }

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	userDB := database.NewUser(db)

	ProductHandler := handlers.NewProductHandler(productDB)
	UserHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn",configs.JwtExperesIn))



	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Authenticator)

		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Post("/", ProductHandler.CreateProduct)
		r.Get("/{id}", ProductHandler.GetProduct)
		r.Get("/", ProductHandler.GetProducts)
		r.Put("/{id}", ProductHandler.UpdateProduct)
		r.Delete("/{id}", ProductHandler.DeleteProduct)
	})
	

	r.Post("/users",UserHandler.Create)
	r.Post("/users/generate_token", UserHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	
	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}