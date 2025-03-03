package initial

import (
	"github.com/abulo/ratel/v3/core/env"
	"github.com/abulo/ratel/v3/core/trace"
	"github.com/spf13/cast"
)

// InitTrace ...
func (initial *Initial) InitTrace() {
	// opt := jaeger.NewJaeger()
	// opt.EnableRPCMetrics = cast.ToBool(initial.Config.Bool("Trace.EnableRPCMetrics"))
	// opt.LocalAgentHostPort = cast.ToString(initial.Config.String("Trace.LocalAgentHostPort"))
	// opt.LogSpans = cast.ToBool(initial.Config.Bool("Trace.LogSpans"))
	// opt.Param = cast.ToFloat64(initial.Config.Float("Trace.Param"))
	// opt.PanicOnError = cast.ToBool(initial.Config.Bool("Trace.PanicOnError"))
	// client := opt.Build().Build()
	// trace.SetGlobalTracer(client)

	trace.SetGlobalTracer(trace.NewConfig(
		trace.WithEndpoint(cast.ToString(initial.Config.String("Trace.LocalAgentHostPort"))),
		trace.WithSampler(1),
		trace.WithName(env.Name()),
	).Build())
}
