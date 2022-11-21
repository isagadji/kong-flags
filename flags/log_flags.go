package flags

import (
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

type LogFlags struct {
	LogLevel string `kong:"optional,name=log-level,env=LOG_LEVEL,default=info"`
}

func (f LogFlags) Init(serviceName string) zerolog.Logger {
	return httplog.NewLogger(serviceName, httplog.Options{
		LogLevel: f.LogLevel,
		JSON:     true,
	})
}
