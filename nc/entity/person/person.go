package person

import (
	"nc/logger"
)

const Name = "name"

type Person struct {
	name string
}

func (p *Person) changeName(newName string) {
	p.name = newName
}

func (p *Person) Process(config map[string]string, logger logger.ILogger) {
	logger.Log("Processing person:", p.name)

	if name, ok := config[Name]; ok {
		logger.Logf("Changing name from %s to %s\n", p.name, name)
		p.changeName(name)
		logger.Logf("Name changed to %s\n", p.name)
	}

	logger.Log("Finished processing person:", p.name)
}
