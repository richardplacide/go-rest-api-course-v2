package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorFetchingComment  = errors.New("error fetching comment by id")
	ErrorNotImpletemented = errors.New("error not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(ctx context.Context, id string) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Getting comment with id: ", id)
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println("Error getting comment: ", err)
		return Comment{}, ErrorFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, body string) error {
	return ErrorNotImpletemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrorNotImpletemented
}

func (s *Service) CreateComment(ctx context.Context, slug, body, author string) (Comment, error) {
	return Comment{}, ErrorNotImpletemented
}
