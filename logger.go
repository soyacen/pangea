package pangea

type (
	Lvl uint8

	JSON interface{}
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	OFF
	panicLevel
	fatalLevel
)

type Logger interface {
	Level() Lvl
	SetLevel(v Lvl)

	Print(i ...interface{})
	Printf(format string, args ...interface{})
	Printj(j JSON)
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Debugj(j JSON)
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Infoj(j JSON)
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Warnj(j JSON)
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Errorj(j JSON)
	Fatal(i ...interface{})
	Fatalj(j JSON)
	Fatalf(format string, args ...interface{})
	Panic(i ...interface{})
	Panicj(j JSON)
	Panicf(format string, args ...interface{})
}
