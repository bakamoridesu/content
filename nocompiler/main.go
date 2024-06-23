package main

func main() {

	logger := Logger{}

	processPerson := func(person *Person) error {
		logger.Log("starting processing person:", person.name)

		config := map[string]string{
			"name": "Soldier",
		}
		person.Process(config, logger)
		// wrappers
		wrappers := []Wrapper{}
		storageWrapper := NewStorageWrapper(person, logger)
		wrappers = append(wrappers, storageWrapper)

		// other wrappers
		// ...

		for _, wrapper := range wrappers {
			if err := wrapper.Process(); err != nil {
				return err
			}
		}
		return nil
	}
	personManager := NewEntityManager(processPerson)
	personManager.RegisterEntity(&Person{"James", Male})
	personManager.RegisterEntity(&Person{"Jane", Female})
	personManager.RegisterEntity(&Person{"John", Other})
	personManager.RegisterEntity(&Person{"Jack", Male})

	persons := personManager.GetEntities()
	for _, person := range persons {
		logger.Log(person.name)
	}
}
