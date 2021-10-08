package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"skrshop/internal/conf"
	"skrshop/internal/data/ent"
)

var ProviderSet = wire.NewSet(NewData, NewBannerRepo, NewDB, NewThemeRepo)

// Data .
type Data struct {
	db  *ent.Client
	log log.Logger
}

func NewDB(conf *conf.Data) *ent.Client {

	opt := ent.Debug()
	client, err := ent.Open(conf.Database.Driver, conf.Database.Source, opt)
	if err != nil {
		panic(err)
	}

	return client
}

// NewData .
func NewData(c *conf.Data, client *ent.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		if err := client.Close(); err != nil {
			log.NewHelper(logger).Error(err)
		}
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{client, logger}, cleanup, nil
}
