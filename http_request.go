package cc_go

import (
	"github.com/crosect/cc-go/web/listener"
	"github.com/crosect/cc-go/web/properties"
	"go.uber.org/fx"
)

func HttpRequestLogOpt() fx.Option {
	return fx.Options(
		ProvideEventListener(listener.NewRequestCompletedLogListener),
		ProvideProps(properties.NewHttpRequestLogProperties),
	)
}
