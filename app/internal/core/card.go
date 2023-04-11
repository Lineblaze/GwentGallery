package core

type repository interface {
	Create
	GetAll
	GetById
	Update
	UpdatePartially
	Delete
}
