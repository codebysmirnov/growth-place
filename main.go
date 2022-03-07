package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	userRepo := repository.NewUserRepo(db)
	userService := user.NewUserService(userRepo)
	userHandlers := handlers.NewUserHandlers(userService)

	router := mux.NewRouter()

	router.HandleFunc("/user", userHandlers.UserCreate).Methods(http.MethodPost)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port), router)
	if err != nil {
		log.Fatal(err)
	}
}
