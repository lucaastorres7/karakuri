package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lucaastorres7/karakuri/internal/config"
	"github.com/lucaastorres7/karakuri/internal/domain/user"
	"github.com/lucaastorres7/karakuri/internal/http/handler"
	"github.com/lucaastorres7/karakuri/internal/repository/postgres"
)

type apiServer struct {
	addr int
	db *pgxpool.Pool
}

func newAPIServer(addr int, db *pgxpool.Pool) *apiServer {
	return &apiServer{
		addr: addr,
		db: db,
	}
}

func (s *apiServer) run() error {
	mux := http.NewServeMux()
	subrouterV1 := http.NewServeMux()
	subrouterV1.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	// Repository Layer
	userRepo := postgres.NewUserRepository(s.db)

	// Service Layer
	userSvc := user.NewService(userRepo)

	// Transport Layer
	userHandler := handler.NewUserHandler(userSvc)

	// Routes
	userHandler.RegisterRoutes(subrouterV1)

	server := http.Server{
		Addr: fmt.Sprintf(":%d", s.addr),
		Handler: subrouterV1,
	}

	log.Println("Server listening on port", s.addr)
	return server.ListenAndServe()
}

func main() {
	ctx := context.Background()
	environment := config.InitConfig()

	dbPool, err := postgres.NewPool(ctx, "")
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	defer dbPool.Close()
	log.Println("Successfully connected to the database!")

	server := newAPIServer(environment.Port, dbPool)
	server.run()
}