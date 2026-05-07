package main

import (
	"flag"
	"os"

	"github.com/lanpang/kratos-layout/internal/pkg/zap"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagConf is the config flag.
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()

	bc, validator, err := loadConfig()
	if err != nil {
		panic(err)
	}

	logger := provideLogger(zap.NewLoggerWithLumberjack(bc.Log))
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger, validator)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
