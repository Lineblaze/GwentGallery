package core

import (
	"context"
	"errors"
	"github.com/Lineblaze/GwentGallery/app/internal/domain"
	"github.com/Lineblaze/GwentGallery/app/internal/domain/errcore"
	"github.com/Lineblaze/GwentGallery/app/internal/repository/database"
)

func (c *core) CreateCard(ctx context.Context, req *domain.CreateCardRequest) (*domain.CreateCardResponse, error) {
	_, err := c.GetUser(ctx, &domain.GetUserRequest{UserId: req.AuthorId})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	req.Status = domain.CardStatusOnChecking

	id, err := c.repo.CreateCard(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetCard(ctx, &domain.GetCardRequest{CardId: id})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.CreateCardResponse{Card: response.Card}, nil
}

func (c *core) GetCard(ctx context.Context, req *domain.GetCardRequest) (*domain.GetCardResponse, error) {
	card, err := c.repo.GetCard(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.CardNotFoundError
		}

		return nil, err
	}

	return &domain.GetCardResponse{Card: card}, nil
}

/*func (c *core) GetAllCards(ctx context.Context, req *domain.GetAllCardsRequest) (*domain.GetAllCardsResponse, error) {
	cards, err := c.repo.GetAllCards(ctx, req)
	if err != nil {
		if errors.Is(err, database.ErrObjectNotFound) {
			return nil, errcore.CardsNotFoundError
		}

		return nil, err
	}

	return &domain.GetAllCardsResponse{Cards: cards}, nil
}*/

func (c *core) UpdateCard(ctx context.Context, req *domain.UpdateCardRequest) (*domain.UpdateCardResponse, error) {
	_, err := c.GetCard(ctx, &domain.GetCardRequest{CardId: req.CardId})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.UpdateCard(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetCard(ctx, &domain.GetCardRequest{CardId: req.CardId})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.UpdateCardResponse{Card: response.Card}, nil
}

/*func (c *core) UpdatePartiallyCard(ctx context.Context, req *domain.UpdatePartiallyCardRequest) (*domain.UpdatePartiallyCardResponse, error) {
	_, err := c.GetCard(ctx, &domain.GetCardRequest{CardId: req.CardId})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.UpdatePartiallyCard(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	response, err := c.GetCard(ctx, &domain.GetCardRequest{CardId: req.CardId})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.UpdatePartiallyCardResponse{Card: response.Card}, nil
}*/

func (c *core) DeleteCard(ctx context.Context, req *domain.DeleteCardRequest) (*domain.DeleteCardResponse, error) {
	_, err := c.GetCard(ctx, &domain.GetCardRequest{CardId: req.CardId})
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	err = c.repo.DeleteCard(ctx, req)
	if err != nil {
		return nil, errcore.NewInternalError(err)
	}

	return &domain.DeleteCardResponse{}, nil
}
