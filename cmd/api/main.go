package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lucaastorres7/karakuri/internal/config"
)

type apiServer struct {
	addr int
}

func newAPIServer(addr int) *apiServer {
	return &apiServer{
		addr: addr,
	}
}

func (s *apiServer) run() error {
	mux := http.NewServeMux()
	subrouterV1 := http.NewServeMux()
	subrouterV1.Handle("/api/v1/", http.StripPrefix("/api/v1", mux))

	server := http.Server{
		Addr: fmt.Sprintf(":%d", s.addr),
		Handler: subrouterV1,
	}

	log.Println("Server listening on port", s.addr)
	return server.ListenAndServe()
}

func main() {
	environment := config.InitConfig()

	server := newAPIServer(environment.Port)
	server.run()
}