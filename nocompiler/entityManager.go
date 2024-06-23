package main

type Processable interface {
	Process(config map[string]string, logger Logger)
}

type EntityManager[T Processable] struct {
	process  func(T) error
	entities []T
}

func NewEntityManager[T Processable](process func(T) error) *EntityManager[T] {
	return &EntityManager[T]{
		process: process,
	}
}

func (p *EntityManager[T]) RegisterEntity(entity T) {

	p.entities = append(p.entities, entity)

	p.process(entity)
}

func (p *EntityManager[T]) GetEntities() []T {
	return p.entities
}
