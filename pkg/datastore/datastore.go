package datastore

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/log"
)

// SmartDS : represnets a datastore which support query feature
// eg : mongo
type SmartDS interface{}

// NewSmartDS creates a new smart datastore
// currently mongo-db is used
func NewSmartDS(reqID int64, conf config.DatastoreConf, l log.Logger) (SmartDS, error) {
	return newMongoDS(reqID, conf, l)
}
