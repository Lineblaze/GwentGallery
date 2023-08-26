package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/Lineblaze/GwentGallery/app/internal/domain"
	"github.com/Lineblaze/GwentGallery/app/internal/repository/database"
	"github.com/jackc/pgx/v5"
)

func (d *driver) CreateCard(ctx context.Context, req *domain.CreateCardRequest) (int64, error) {
	row := d.conn.QueryRow(ctx, `insert into cards(name, category_id, description, image_id, provision, power, faction, rarity, card_type, color, set, image_author, author_id, status)
                                     values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
                                     returning id`,
		req.Name,
		req.CategoryId,
		req.Description,
		req.ImageId,
		req.Provision,
		req.Power,
		req.Faction,
		req.Rarity,
		req.CardType,
		req.Color,
		req.Set,
		req.ImageAuthor,
		req.AuthorId,
		req.Status,
	)
	var cardId int64

	if err := row.Scan(&cardId); err != nil {
		return 0, fmt.Errorf("%w: scanning Card id", err)
	}

	return cardId, nil
}

func (d *driver) GetCard(ctx context.Context, req *domain.GetCardRequest) (*domain.Card, error) {
	row := d.conn.QueryRow(ctx, `select 
                                       c.id, c.name, c.category_id, c.description, c.image_id, c.provision, c.power,
                                       c.faction, c.rarity, c.card_type, c.color, c.set, c.image_author, c.author_id, c.status,
                                       c.created_at, c.updated_at
                                     from cards c
                                     where c.id = $1`, req.CardId)

	var card domain.Card
	if err := row.Scan(
		&card.Id,
		&card.Name,
		&card.CategoryId,
		&card.Description,
		&card.ImageId,
		&card.Provision,
		&card.Power,
		&card.Faction,
		&card.Rarity,
		&card.CardType,
		&card.Color,
		&card.Set,
		&card.ImageAuthor,
		&card.AuthorId,
		&card.Status,
		&card.CreatedAt,
		&card.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		return nil, fmt.Errorf("%w: scanning card", err)
	}

	return &card, nil
}

/*
	func (d *driver) GetAllCards(ctx context.Context, req *domain.GetAllCardsRequest) ([]domain.Card, error) {
		rows, err := d.conn.Query(ctx,
			`SELECT *
				 FROM cards`)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, database.ErrObjectNotFound
		}
		defer rows.Close()
		var cards []domain.Card
		for rows.Next() {
			card := domain.Card{}
			err = rows.Scan(
				&card.Id,
				&card.Name,
				&card.CategoryId,
				&card.Description,
				&card.ImageId,
				&card.Provision,
				&card.Power,
				&card.Faction,
				&card.Rarity,
				&card.CardType,
				&card.Color,
				&card.Set,
				&card.ImageAuthor,
				&card.AuthorId,
			)
			cards = append(cards, card)
		}
		return cards, nil

}
*/
func (d *driver) UpdateCard(ctx context.Context, req *domain.UpdateCardRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `update cards 
                                     set name = $2, category_id = $3, description = $4, provision = $5, power = $6, 
                                		 faction = $7, rarity = $8, card_type = $9, color = $10, updated_at = now()
                                     where id = $1
                                     returning true`,
		req.CardId,
		req.Name,
		req.CategoryId,
		req.Description,
		req.Provision,
		req.Power,
		req.Faction,
		req.Rarity,
		req.CardType,
		req.Color,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("updating card: %w", err)
	}

	return nil
}

/*func (d *driver) UpdatePartiallyCard(ctx context.Context, req *domain.UpdatePartiallyCardRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `update cards
                                     set name = $2, category_id = $3, description = $4, provision = $5, power = $6,
                                		 faction = $7, rarity = $8, card_type = $9, color = $10, updated_at = now()
                                     where id = $1
                                     returning true`,
		req.CardId,
		req.Name,
		req.CategoryId,
		req.Description,
		req.Provision,
		req.Power,
		req.Faction,
		req.Rarity,
		req.CardType,
		req.Color,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("partially updating card: %w", err)
	}

	return nil
}*/

func (d *driver) DeleteCard(ctx context.Context, req *domain.DeleteCardRequest) error {
	var ok bool
	err := d.conn.QueryRow(ctx, `delete from cards 
                                     where id = $1
                                     returning true`,
		req.CardId,
	).Scan(&ok)
	if err != nil {
		return fmt.Errorf("deleting card: %w", err)
	}

	return nil
}
