package repository

type InmemRepo interface{}

type inmemrepo struct{}

func NewInmemRepository() InmemRepo {
	return &inmemrepo{}
}
