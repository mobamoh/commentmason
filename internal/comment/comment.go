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
	CreateComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
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

func (service *Service) UpdateComment(ctx context.Context, id string, comment Comment) (Comment, error) {
	updatedCmt, err := service.Store.UpdateComment(ctx, id, comment)
	if err != nil {
		return Comment{}, err
	}
	return updatedCmt, nil
}

func (service *Service) DeleteComment(ctx context.Context, id string) error {
	return service.Store.DeleteComment(ctx, id)
}

func (service *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	insertedCmt, err := service.Store.CreateComment(ctx, comment)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
