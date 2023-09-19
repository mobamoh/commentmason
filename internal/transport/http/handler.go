package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Server  *http.Server
	Service CommentService
}

func NewHandler(service CommentService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.PutComment).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.DeleteComment).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutdown gracefully")
	return nil
}
