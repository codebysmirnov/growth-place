package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "growth-place/docs"
	"growth-place/middlewares"

	"growth-place/config"
	"growth-place/src/handlers"
	"growth-place/src/repository"
	"growth-place/src/services/user"
)

// @title Growth-place API
// @version 1.0
// @description This is a growth service for managing personal targets
// @license.name MIT
// @BasePath /
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = cfg.Validate()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
			cfg.Database.Host,
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Name,
			cfg.Database.Port,
			cfg.Database.SSLMode,
		)),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}

	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	userRepo := repository.NewUserRepo(db)
	userService := user.NewUserService(userRepo, logger, cfg.JWT.Secret)
	userHandlers := handlers.NewUserHandlers(userService)

	// without authorization
	router := mux.NewRouter()
	v1 := router.PathPrefix("/v1/").Subrouter()
	v1.HandleFunc("/user", userHandlers.UserCreate).Methods(http.MethodPost)
	v1.HandleFunc("/user/authorization", userHandlers.UserAuthorization).Methods(http.MethodPost)

	// authorized routes
	v1Auth := router.PathPrefix("/v1/").Subrouter()
	authMiddleware := middlewares.NewAuthorization(cfg.JWT.Secret, logger)
	v1Auth.Use(authMiddleware)
	v1Auth.HandleFunc("/user", userHandlers.Profile).Methods(http.MethodGet)
	v1Auth.HandleFunc("/user", userHandlers.Delete).Methods(http.MethodDelete)
	v1Auth.HandleFunc("/user/password", userHandlers.PasswordEdit).Methods(http.MethodPost)

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), router)
	if err != nil {
		log.Fatal(err)
	}
}
