package service

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"

	"github.com/alexjch/queryrepo/internal/cmd/signals"
	"github.com/alexjch/queryrepo/internal/handlers"
)

type Service struct{}

func (s *Service) Start(port int, repoUrl *url.URL) {

	router := mux.NewRouter()
	router.StrictSlash(false)
	subRouterV1 := router.PathPrefix("/api/v1").Subrouter()
	subRouterV1.StrictSlash(false)

	portStr := fmt.Sprintf(":%d", port)

	/* Endpoints definition  */
	rootHandlerV1 := handlers.NewRootHandlerV1(repoUrl)
	subRouterV1.HandleFunc("/", rootHandlerV1).Methods("POST")
	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	router.HandleFunc("/readiness", handlers.ReadinessHandler).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         portStr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	/* Start the server */
	go func() {
		log.Println("Starting Server on port:", portStr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	signals.WaitForShutdown(server)
}
