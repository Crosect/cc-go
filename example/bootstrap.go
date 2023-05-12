package example

// ==================================================
// ==== Example about how to bootstrap your app =====
// ==================================================

import (
	"github.com/crosect/cc-go"
	"github.com/crosect/cc-go/pubsub"
	"github.com/crosect/cc-go/pubsub/executor"
	"go.uber.org/fx"
)

func All() fx.Option {
	return fx.Options(
		// Required
		ccgo.AppOpt(),
		ccgo.PropertiesOpt(),

		// When you want to enable event publisher
		ccgo.EventOpt(),
		// When you want handle event in simple synchronous way
		ccgo.SupplyEventBusOpt(pubsub.WithEventExecutor(executor.NewSyncExecutor())),
		// Or want a custom executor, such as using worker pool
		fx.Provide(NewSampleEventExecutor),
		ccgo.ProvideEventBusOpt(func(executor *SampleEventExecutor) pubsub.EventBusOpt {
			return pubsub.WithEventExecutor(executor)
		}),

		// When you want to use default logging strategy.
		ccgo.LoggingOpt(),
		// When you want to enable http request log
		ccgo.HttpRequestLogOpt(),

		// When you want to enable actuator endpoints.
		// By default, we provide HealthService and InfoService.
		ccgo.ActuatorEndpointOpt(),
		// When you want to provide build info to above InfoService.
		ccgo.BuildInfoOpt(Version, CommitHash, BuildTime),
		// When you want to provide custom health checker and informer
		ccgo.ProvideHealthChecker(NewSampleHealthChecker),
		ccgo.ProvideInformer(NewSampleInformer),

		// When you want to enable http client auto config with contextual client by default
		ccgo.HttpClientOpt(),

		// When you want to tell cc-go to load your properties.
		ccgo.ProvideProps(NewSampleProperties),

		// When you want to declare a service
		fx.Provide(NewSampleService),

		// When you want to register your event listener.
		ccgo.ProvideEventListener(NewSampleListener),

		// Graceful shutdown.
		// OnStop hooks will run in reverse order.
		ccgo.OnStopEventOpt(),
	)
}
