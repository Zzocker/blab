package adapters

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/internal/logger"
	"github.com/Zzocker/blab/pkg/datastore"
)

const (
	adaptersPkg = "adapters"
)

func CreateUserstore(conf *config.DatastoreConf) (*userstore, error) {
	logger.L.Infof(-1, adaptersPkg, "createing user store")
	ds, err := datastore.NewSmartDS(-1, *conf, logger.L)
	if err != nil {
		return nil, err
	}
	logger.L.Infof(-1, adaptersPkg, "userstore created")
	return &userstore{
		ds: ds,
	}, nil
}
