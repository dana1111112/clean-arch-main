package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/configs"
	"architecture_go/services/contact/internal/delivery/http"
	postgres2 "architecture_go/services/contact/internal/repository/storage/postgres"
	"architecture_go/services/contact/internal/useCase/contact"
	"architecture_go/services/contact/internal/useCase/group"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	storage, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer storage.Close()

	log.Println("Successfully connected to database!")

	contactRepo := postgres2.NewContactRepository(storage)
	contactUseCase := contact.NewContactUseCase(&contactRepo)
	groupUseCase := group.NewGroupUseCase(&contactRepo, &contactRepo)
	delivery := http.NewContactHTTP(contactUseCase, groupUseCase)
	delivery.Run(cfg)
}
