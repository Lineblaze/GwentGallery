package errcore

var (
	CardNotFoundError = &CoreError{
		Message: "Card not found.",
		Code:    "Card.not_found",
		Type:    NotFoundType,
	}

	CardsNotFoundError = &CoreError{
		Message: "Cards not found.",
		Code:    "Cards.not_found",
		Type:    NotFoundType,
	}
)
