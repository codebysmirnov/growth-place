package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"growth-place/application/handlers"
	"growth-place/application/repository"
	"growth-place/application/services/user"
	"growth-place/config"
)

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

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	userRepo := repository.NewUserRepo(db)
	userService := user.NewUserService(userRepo, logger, cfg.JWT.Secret)
	userHandlers := handlers.NewUserHandlers(userService)

	router := mux.NewRouter()

	router.HandleFunc("/user", userHandlers.UserCreate).Methods(http.MethodPost)
	router.HandleFunc("/user/password", userHandlers.UserPasswordEdit).Methods(http.MethodPost)
	router.HandleFunc("/user/authorization", userHandlers.UserAuthorization).Methods(http.MethodPost)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), router)
	if err != nil {
		log.Fatal(err)
	}
}
