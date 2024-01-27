package gps

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/google/wire"
)

func ProvideLogger() watermill.LoggerAdapter {
	return watermill.NewStdLogger(true, true)
}

var Set = wire.NewSet(ProvideLogger)
