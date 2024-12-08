package middleware

import "automatic-doodle/types"

type AuthenticationService interface {
	GetUserByToken(*string) (types.AuthenticatedUser, error)
}
type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}
