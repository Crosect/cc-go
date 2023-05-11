package cc_go

import (
	"github.com/crosect/cc-go/build"
	"go.uber.org/fx"
)

func BuildInfoOpt(version string, commitHash string, time string) fx.Option {
	return fx.Options(
		fx.Supply(build.Version(version)),
		fx.Supply(build.CommitHash(commitHash)),
		fx.Supply(build.Time(time)),
		ProvideInformer(build.NewInformer),
	)
}
