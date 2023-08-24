package domain

import "time"

const (
	_ Faction = iota
	CardFactionNeutral
	CardFactionMonsters
	CardFactionNilfgaard
	CardFactionNorthearnRealms
	CardFactionScoiatael
	CardFactionScellige
	CardFactionSyndicate
)

var CardFactions = map[Faction]string{
	CardFactionNeutral:         "NEUTRAL",
	CardFactionMonsters:        "MONSTERS",
	CardFactionNilfgaard:       "NILFAGAARD",
	CardFactionNorthearnRealms: "NORTHEARN REALMS",
	CardFactionScoiatael:       "SCOIA'TAEL",
	CardFactionScellige:        "SCELLIGE",
	CardFactionSyndicate:       "SYNDICATE",
}

const (
	_ Rarity = iota
	CardRarityCommon
	CardRarityRare
	CardRarityEpic
	CardRarityLegendary
)

var CardRarities = map[Rarity]string{
	CardRarityCommon:    "COMMON",
	CardRarityRare:      "RARE",
	CardRarityEpic:      "EPIC",
	CardRarityLegendary: "LEGENDARY",
}

const (
	_ CardType = iota
	CardTypeUnit
	CardTypeSpecial
	CardTypeArtifact
	CardTypeUndefined
)

var CardTypes = map[CardType]string{
	CardTypeUnit:      "UNIT",
	CardTypeSpecial:   "SPECIAL",
	CardTypeArtifact:  "ARTIFACT",
	CardTypeUndefined: "UNDEFINED",
}

const (
	_ Color = iota
	CardColorGold
	CardColorBronze
)

var CardColors = map[Color]string{
	CardColorGold:   "GOLD",
	CardColorBronze: "BRONZE",
}

const (
	_ Set = iota
	CardSetStarter
	CardSetBase
	CardSetCrimsonCurse
	CardSetNovigrad
	CardSetIronJudgment
	CardSetMerchantsOfOfir
	CardSetMasterMirror
	CardSetWayOfTheWitcher
	CardSetPriceOfPower
)

var CardSets = map[Set]string{
	CardSetStarter:         "STARTER",
	CardSetBase:            "BASE",
	CardSetCrimsonCurse:    "CRIMSON CURSE",
	CardSetNovigrad:        "NOVIGRAD",
	CardSetIronJudgment:    "IRON JUDGMENT",
	CardSetMerchantsOfOfir: "MERCHANTS OF OFIR",
	CardSetMasterMirror:    "MASTER MIRROR",
	CardSetWayOfTheWitcher: "WAY OF THE WITCHER",
	CardSetPriceOfPower:    "PRICE OF POWER",
}

const (
	_ Status = iota
	CardStatusUnChecking
	CardStatusOnChecking
	CardStatusApproved
	CardStatusDeclined
)

type (
	Faction  int8
	Rarity   int8
	CardType int8
	Color    int8
	Set      int8
	Status   int8

	Card struct {
		Id           int64      `json:"id"`
		Name         string     `json:"name"`
		CategoryId   *uint8     `json:"category_id,omitempty"`
		Description  string     `json:"description"`
		ImageId      string     `json:"image_id"`
		Provision    uint8      `json:"provision"`
		Power        *uint8     `json:"power,omitempty"`
		Faction      Faction    `json:"faction"`
		FactionName  string     `json:"faction_name"`
		Rarity       Rarity     `json:"rarity"`
		RarityName   string     `json:"rarity_name"`
		CardType     CardType   `json:"card_type"`
		CardTypeName string     `json:"card_type_name"`
		Color        Color      `json:"color"`
		ColorName    string     `json:"color_name"`
		Set          Set        `json:"set"`
		SetName      string     `json:"set_name"`
		ImageAuthor  string     `json:"image_author"`
		AuthorId     int64      `json:"author_id"`
		Status       Status     `json:"status"`
		CreatedAt    time.Time  `json:"created_at"`
		UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	}

	CreateCardRequest struct {
		Name        string   `json:"name"`
		CategoryId  *uint8   `json:"category_id,omitempty"`
		Description string   `json:"description"`
		ImageId     string   `json:"image_id"`
		Provision   uint8    `json:"provision"`
		Power       *uint8   `json:"power,omitempty"`
		Faction     Faction  `json:"faction"`
		Rarity      Rarity   `json:"rarity"`
		CardType    CardType `json:"card_type"`
		Color       Color    `json:"color"`
		Set         Set      `json:"card_set"`
		ImageAuthor string   `json:"image_author"`
		AuthorId    int64    `json:"author_id"`
		Status      Status   `json:"status"`
	}

	CreateCardResponse struct {
		Card *Card `json:"data"`
	}

	GetCardRequest struct {
		CardId int64 `json:"-"`
	}
	GetCardResponse struct {
		Card *Card `json:"data"`
	}
	GetAllCardsRequest struct {
		//Search    string
		CategoryId  *uint8   `json:"category_id,omitempty"`
		Provision   uint8    `json:"provision"`
		Power       *uint8   `json:"power,omitempty"`
		Faction     Faction  `json:"faction"`
		Rarity      Rarity   `json:"rarity"`
		CardType    CardType `json:"card_type"`
		Color       Color    `json:"color"`
		Set         Set      `json:"card_set"`
		ImageAuthor string   `json:"image_author"`
		AuthorId    *int64   `json:"author_id"`
		Status      Status   `json:"status"`
	}
	GetAllCardsResponse struct {
		Cards []Card `json:"data"`
	}
	UpdateCardRequest struct {
		CardId      int64    `json:"-"`
		Name        string   `json:"name"`
		CategoryId  *uint8   `json:"category_id,omitempty"`
		Description string   `json:"description"`
		Provision   uint8    `json:"provision"`
		Power       *uint8   `json:"power,omitempty"`
		Faction     Faction  `json:"faction"`
		Rarity      Rarity   `json:"rarity"`
		CardType    CardType `json:"card_type"`
		Color       Color    `json:"color"`
	}
	UpdateCardResponse struct {
		Card *Card `json:"data"`
	}
	/*
		UpdatePartiallyCardRequest struct {
			CardId      int64    `json:"-"`
			Name        *string  `json:"name"`
			CategoryId  *uint8   `json:"category_id,omitempty"`
			Description *string  `json:"description"`
			Provision   *uint8   `json:"provision"`
			Power       *uint8   `json:"power,omitempty"`
			Faction     *Faction `json:"faction"`
			Rarity      *Rarity  `json:"rarity"`
			CardType    CardType `json:"card_type"`
			Color       *Color   `json:"color"`
		}
	*/
	UpdatePartiallyCardResponse struct {
		Card *Card `json:"data"`
	}
	DeleteCardRequest struct {
		CardId int64 `json:"card_id"`
	}
	DeleteCardResponse struct {
	}
)
