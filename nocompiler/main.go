package main

import (
	"nocompiler/entity/person"
	"nocompiler/logger"
	"nocompiler/manager"
	"nocompiler/wrapper"
	"nocompiler/wrapper/storage"
)

func main() {

	logger := logger.NewLogger()

	processPerson := func(p *person.Person) error {
		logger.Log("starting processing person:", p.Name)

		config := map[string]string{
			"name": "Soldier",
		}
		p.Process(config, logger)
		// other runners
		wrappers := []wrapper.Wrapper{}
		storageWrapper := storage.New(p, logger)
		wrappers = append(wrappers, storageWrapper)

		// other wrappers
		// ...

		for _, wrapper := range wrappers {
			if err := wrapper.Run(); err != nil {
				return err
			}
		}
		return nil
	}

	personManager := manager.New(processPerson)

	personManager.RegisterEntity(&person.Person{Name: "James", Gender: person.Male})
	personManager.RegisterEntity(&person.Person{Name: "John", Gender: person.Other})
	personManager.RegisterEntity(&person.Person{Name: "Jane", Gender: person.Female})
	personManager.RegisterEntity(&person.Person{Name: "Jim", Gender: person.Male})

	persons := personManager.GetEntities()
	for _, person := range persons {
		logger.Log(person.Name)
	}
}
