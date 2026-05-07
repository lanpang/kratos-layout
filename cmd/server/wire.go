//go:build wireinject
// +build wireinject

package main

import (
	"buf.build/go/protovalidate"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tpl-x/kratos/internal/biz"
	"github.com/tpl-x/kratos/internal/conf"
	"github.com/tpl-x/kratos/internal/data"
	"github.com/tpl-x/kratos/internal/server"
	"github.com/tpl-x/kratos/internal/service"
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
