package ccgo

import (
	"github.com/crosect/cc-go/log"
	webLog "github.com/crosect/cc-go/web/log"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

func LoggingOpt() fx.Option {
	return fx.Options(
		ProvideProps(log.NewProperties),
		fx.Provide(NewLogger),
		fx.Invoke(RegisterLogger),
	)
}

type NewLoggerOut struct {
	fx.Out
	Core log.Logger
	Web  log.Logger `name:"web_logger"`
}

func NewLogger(props *log.Properties) (NewLoggerOut, error) {
	out := NewLoggerOut{}
	// Create new logger instance
	logger, err := log.NewDefaultLogger(&log.Options{
		Development:    props.Development,
		JsonOutputMode: props.JsonOutputMode,
		CallerSkip:     props.CallerSkip,
	})
	if err != nil {
		return out, errors.WithMessage(err, "init logger failed")
	}
	out.Core = logger
	out.Web = logger.Clone(log.AddCallerSkip(1))
	return out, nil
}

type RegisterLoggerIn struct {
	fx.In
	Core log.Logger
	Web  log.Logger `name:"web_logger"`
}

func RegisterLogger(in RegisterLoggerIn) {
	log.ReplaceGlobal(in.Core)
	webLog.ReplaceGlobal(in.Web)
}
