package server

import (
	"blockchain/internal/config"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	cfg    *config.Config
	router *gin.Engine
	db     *pgxpool.Pool
}

func NewServer(cfg *config.Config, router *gin.Engine, db *pgxpool.Pool) *Server {
	return &Server{
		cfg:    cfg,
		router: router,
		db:     db,
	}
}

func (s *Server) Start() {
	// Issuance waits for blockchain confirmation and can legitimately
	// take well over 10 seconds before the handler writes a response.
	const (
		readTimeout       = 30 * time.Second
		writeTimeout      = 3 * time.Minute
		idleTimeout       = 60 * time.Second
		readHeaderTimeout = 5 * time.Second
	)

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           s.router,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	log.Println("server running on :8080")

	if err := httpServer.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {

		log.Fatal(err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("shutting down server")

	s.db.Close()

	return nil
}
