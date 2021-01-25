package log

import "go.uber.org/zap"

type logger struct {
	*zap.SugaredLogger
}