package logger

type ILoggerProvider interface {
	Error(namespacem, message string)
	Info(namespacem, message string)
}
