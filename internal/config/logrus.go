package config

import (
	"github.com/lamaleka/boilerplate-golang/internal/entity"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *entity.ConfLog) *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.Level(config.Level))
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	return log
}
