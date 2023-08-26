package database

import (
	"context"
	"github.com/Lineblaze/GwentGallery/app/internal/domain"
)

type Database interface {
	CreateCard(ctx context.Context, req *domain.CreateCardRequest) (int64, error)
	GetCard(ctx context.Context, req *domain.GetCardRequest) (*domain.Card, error)
	/*	GetAllCards(ctx context.Context, req *domain.GetAllCardsRequest) ([]domain.Card, error)*/
	UpdateCard(ctx context.Context, req *domain.UpdateCardRequest) error
	/*	UpdatePartiallyCard(ctx context.Context, req *domain.UpdatePartiallyCardRequest) error*/
	DeleteCard(ctx context.Context, req *domain.DeleteCardRequest) error

	GetUser(ctx context.Context, req *domain.GetUserRequest) (*domain.User, error)
	/*	GetUsers(ctx context.Context, req *domain.GetUserRequest) ([]domain.User, error)*/
	GetUserByLogin(ctx context.Context, login string) (*domain.User, error)
	CheckUsernameUniqueness(ctx context.Context, username string) error
	CheckEmailUniqueness(ctx context.Context, email string) error
	DeleteUser(ctx context.Context, req *domain.DeleteUserRequest) error

	SignUp(ctx context.Context, req *domain.SignUpRequest) (int64, error)

	Close()
}
