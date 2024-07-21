package main

type store struct {
}

func NewStore() *store {
	return &store{}
}

func (s *store) Register() error {
	return nil
}
