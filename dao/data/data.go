package data

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/LiuKaHo/driverModule/conf"
	"github.com/LiuKaHo/driverModule/dao/ent"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	drv, err := sql.Open(
		c.GetProduct().GetDatabase().GetDriver(),
		c.GetProduct().GetDatabase().GetSource(),
	)
	if err != nil {
		return nil, nil, err
	}
	sqlDrv := dialect.Debug(drv)
	client := ent.NewClient(ent.Driver(sqlDrv))
	return &Data{
		client,
	}, nil, nil
}
