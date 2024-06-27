package storage

import (
	"nocompiler/entity/person"
	"nocompiler/logger"
)

type StorageWrapper struct {
	p *person.Person
	l logger.ILogger
}

func New(p *person.Person, l logger.ILogger) *StorageWrapper {
	return &StorageWrapper{
		p: p,
		l: l,
	}
}

func (s *StorageWrapper) Run() error {
	s.l.Log("Storing", s.p.Name)

	if err := s.Store(); err != nil {
		s.l.Log("Error storing", s.p.Name)
		return err
	}
	s.l.Log("Stored", s.p.Name)
	return nil
}

func (s *StorageWrapper) Store() error {
	// store
	return nil
}
