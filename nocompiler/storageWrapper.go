package main

type StorageWrapper struct {
	p *Person
	l Logger
}

func NewStorageWrapper(p *Person, l Logger) *StorageWrapper {
	return &StorageWrapper{
		p: p,
		l: l,
	}
}

func (s *StorageWrapper) Process() error {
	s.l.Log("Storing", s.p.name)

	if err := s.Store(); err != nil {
		s.l.Log("Error storing", s.p.name)
		return err
	}
	s.l.Log("Stored", s.p.name)
	return nil
}

func (s *StorageWrapper) Store() error {
	// store
	return nil
}
