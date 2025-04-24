package config

import (
	"fmt"
	"time"

	"github.com/lamaleka/boilerplate-golang/internal/entity"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config *entity.ConfDb, log *logrus.Logger) *gorm.DB {

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.Username, config.Password, config.Host, config.Port, config.Name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		}),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	connection.SetMaxIdleConns(config.Pool.Idle)
	connection.SetMaxOpenConns(config.Pool.Max)
	connection.SetConnMaxLifetime(time.Second * time.Duration(config.Pool.Lifetime))

	return db
}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
