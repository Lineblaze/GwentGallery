package core

import (
	"context"
	"fmt"
	"github.com/Lineblaze/GwentGallery/app/internal/domain"
	"github.com/Lineblaze/GwentGallery/app/internal/repository/database"
	"github.com/Lineblaze/GwentGallery/app/internal/repository/database/postgresql"
	"os"
)

// Core represents business logic layer interface.
type Core interface {
	CreateCard(ctx context.Context, req *domain.CreateCardRequest) (*domain.CreateCardResponse, error)
	GetCard(ctx context.Context, req *domain.GetCardRequest) (*domain.GetCardResponse, error)
	/*GetAllCards(ctx context.Context, req *domain.GetAllCardsRequest) (*domain.GetAllCardsResponse, error)*/
	UpdateCard(ctx context.Context, req *domain.UpdateCardRequest) (*domain.UpdateCardResponse, error)
	/*UpdatePartiallyCard(ctx context.Context, req *domain.UpdatePartiallyCardRequest) (*domain.UpdatePartiallyCardResponse, error)*/
	DeleteCard(ctx context.Context, req *domain.DeleteCardRequest) (*domain.DeleteCardResponse, error)

	SignIn(ctx context.Context, req *domain.SignInRequest) (*domain.SignInResponse, error)
	SignUp(ctx context.Context, req *domain.SignUpRequest) (*domain.SignUpResponse, error)

	GetUser(ctx context.Context, req *domain.GetUserRequest) (*domain.GetUserResponse, error)
	/*GetUsers(ctx context.Context, request *domain.GetUsersRequest) (*domain.GetUsersResponse, error)*/
	/*UpdateUser(ctx context.Context, req *domain.UpdateUserRequest) (*domain.UpdateUserResponse, error)*/
	DeleteUser(ctx context.Context, req *domain.DeleteUserRequest) (*domain.DeleteUserResponse, error)
}

// core implements Core interface.
type core struct {
	repo database.Database
}

// New returns Core instance.
func New(ctx context.Context) (Core, error) {
	db, err := postgresql.NewDriver(ctx, os.Getenv("DATABASE_DSN"))
	if err != nil {
		return nil, fmt.Errorf("creating postgresql driver: %w", err)
	}

	return &core{repo: db}, nil
}
