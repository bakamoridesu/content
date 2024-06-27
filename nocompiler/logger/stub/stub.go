package stub

import (
	"nocompiler/logger"
)

type Stub struct{}

func (s *Stub) Log(args ...interface{}) {}

func (s *Stub) Logf(format string, args ...interface{}) {}

func NewStub() logger.ILogger {
	return &Stub{}
}
