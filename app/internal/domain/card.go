package domain

import "time"

type Card struct {
	ID          int64      `mapstructure:"id"`
	Name        string     `mapstructure:"name"`
	Category    string     `mapstructure:"category"`
	Description string     `mapstructure:"description"`
	ImageID     *string    `mapstructure:"image_id"`
	Provision   uint32     `mapstructure:"provisionp"`
	Power       uint32     `mapstructure:"power"`
	Faction     string     `mapstructure:"faction"`
	Rarity      string     `mapstructure:"rarity"`
	Color       string     `mapstructure:"color"`
	CardSet     string     `mapstructure:"card_set"`
	ImageAuthor string     `mapstructure:"image_author"`
	CreatedAt   time.Time  `mapstructure:"created_at"`
	UpdatedAt   *time.Time `mapstructure:"updated_at"`
}
