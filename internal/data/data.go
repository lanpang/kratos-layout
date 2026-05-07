package data

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tpl-x/kratos/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewGreeterRepo,
)

// Data wraps data-layer clients.
type Data struct {
	db *gorm.DB
}

// DB returns the GORM database handle for repositories.
func (d *Data) DB() *gorm.DB {
	return d.db
}

// NewData initializes data-layer resources.
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(log.With(logger, "module", "data"))

	db, err := newGORM(c)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		helper.Info("closing the data resources")
		sqlDB, err := db.DB()
		if err != nil {
			helper.Errorf("get gorm sql db failed: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			helper.Errorf("close gorm sql db failed: %v", err)
		}
	}
	return &Data{db: db}, cleanup, nil
}

func newGORM(c *conf.Data) (*gorm.DB, error) {
	if c == nil || c.Database == nil {
		return nil, fmt.Errorf("data.database config is required")
	}

	switch c.Database.Driver {
	case "postgres", "postgresql", "":
		return gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database driver %q", c.Database.Driver)
	}
}
