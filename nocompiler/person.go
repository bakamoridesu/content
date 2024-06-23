package main

const (
	Male = iota
	Female
	Other
)

const (
	Name = "name"
)

type Person struct {
	name   string
	gender int
}

func (p *Person) changeName(name string, logger Logger) error {
	if p.gender != Male {
		return nil
	}
	logger.Log("Changing name to", name)
	p.name = name
	logger.Log("Changed name to", name)
	return nil
}

func (p *Person) Process(config map[string]string, logger Logger) {
	logger.Log("Processing...", p.name)
	if name, ok := config[Name]; ok {
		err := p.changeName(name, logger)
		if err != nil {
			logger.Log("Error processing", p.name)
			return
		}
	}
	logger.Log("Done processing", p.name)
}
