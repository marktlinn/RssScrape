package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server(port string) {
	srv := &http.Server{
		Handler: newRouter(),
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Server running on port: %s\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen on port %s\n", port)
	}
}
