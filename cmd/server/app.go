package main

import (
	"buf.build/go/protovalidate"
	zapv2 "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/tpl-x/kratos/internal/conf"
)

func loadConfig() (*conf.Bootstrap, protovalidate.Validator, error) {
	validator, err := protovalidate.New()
	if err != nil {
		return nil, nil, err
	}

	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		return nil, nil, err
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		return nil, nil, err
	}

	if err := validator.Validate(&bc); err != nil {
		return nil, nil, err
	}

	return &bc, validator, nil
}

// provideLogger creates a Kratos logger with service information.
func provideLogger(zapLogger *zapv2.Logger) log.Logger {
	return log.With(zapLogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}

// newKratosApp creates the Kratos application.
func newKratosApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(gs, hs),
	)
}
