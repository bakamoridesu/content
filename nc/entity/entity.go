package entity

import "nc/logger"

type Entity interface {
	Process(config map[string]string, logger logger.ILogger)
}
