package middleware

import (
	"github.com/Zzocker/blab/config"
	"github.com/Zzocker/blab/pkg/errors"
	"github.com/Zzocker/blab/pkg/log"
)

type mi interface {
	build(config.C) errors.E
}

var (
	factory = []mi{
		&oauthMi{},
	}
)

func BuildMiddleware(conf config.C) {
	for i := range factory {
		err := factory[i].build(conf)
		if err != nil {
			log.L.Fatal(err)
		}
	}
	log.L.Info("All middleware built")
}
