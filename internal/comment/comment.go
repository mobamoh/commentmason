package comment

import (
	"context"
	"errors"
	"log"
)

var (
	ErrFetchingComment = errors.New("error fetching the comment!")
	ErrNotImplemented  = errors.New("method not implemented!")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (service *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	comment, err := service.Store.GetComment(ctx, id)
	if err != nil {
		log.Printf("comment: %v", err)
		return Comment{}, ErrFetchingComment
	}
	return comment, nil
}

func (Service *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (Service *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (Service *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
