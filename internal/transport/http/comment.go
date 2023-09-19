package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mobamoh/commentmason/internal/comment"
)

type CommentService interface {
	GetComment(context.Context, string) (comment.Comment, error)
	CreateComment(context.Context, comment.Comment) (comment.Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
}

type Response struct {
	message string
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}
	cmt, err := h.Service.CreateComment(r.Context(), cmt)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) PutComment(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := h.Service.DeleteComment(r.Context(), id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(Response{message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
