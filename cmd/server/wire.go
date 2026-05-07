//go:build wireinject
// +build wireinject

package main

import (
	"buf.build/go/protovalidate"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lanpang/kratos-layout/internal/biz"
	"github.com/lanpang/kratos-layout/internal/conf"
	"github.com/lanpang/kratos-layout/internal/data"
	"github.com/lanpang/kratos-layout/internal/server"
	"github.com/lanpang/kratos-layout/internal/service"
)

func wireApp(*conf.Server, *conf.Data, log.Logger, protovalidate.Validator) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newKratosApp,
	))
}
