package flags

type LogFlags struct {
	LogLevel string `kong:"optional,name=log-level,env=LOG_LEVEL,default=info"`
}
