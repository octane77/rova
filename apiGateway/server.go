package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ApplicationServer struct {
	server *http.Server
}

func (a *ApplicationServer) Start() {

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Printf("Listen: %s", err)
			} else {
				log.Printf("could not start server: %v", err)
			}
		}
	}()
	log.Printf("Server is Running on Port: %s", a.server.Addr)

	// GRACEFUL SHUTDOWN
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting Down Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Forced To ShutDown, %v", err)
	}
	log.Println("Server Exiting")
}

func NewApplicationServer(handler *gin.Engine) *ApplicationServer {
	godotenv.Load()
	port := os.Getenv("PORT")
	return &ApplicationServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}
}
