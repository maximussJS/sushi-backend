package logger

type ILogger interface {
	Log(message string)
	Warn(message string)
	Error(message string)
	Trace(message string)
	Debug(message string)
	Fatal(message string)
	Panic(message string)
}
