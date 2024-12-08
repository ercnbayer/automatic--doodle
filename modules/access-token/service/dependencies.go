package service

type ConfigModule interface {
	GetConfig(string) (string, error)
	GetConfigInt(string) (int, error)
}

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}
