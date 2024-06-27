package person

import "nocompiler/logger"

const (
	Male = iota
	Female
	Other
)

const (
	Name = "name"
)

type Person struct {
	Name   string
	Gender int
}

func (p *Person) changeName(name string, logger logger.ILogger) error {
	if p.Gender != Male {
		return nil
	}
	logger.Log("Changing name to", name)
	p.Name = name
	logger.Log("Changed name to", name)
	return nil
}

func (p *Person) Process(config map[string]string, logger logger.ILogger) {
	logger.Log("Processing...", p.Name)
	if name, ok := config[Name]; ok {
		err := p.changeName(name, logger)
		if err != nil {
			logger.Log("Error processing", p.Name)
			return
		}
	}
	logger.Log("Done processing", p.Name)
}
