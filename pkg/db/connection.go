package db

import (
	"endeus/wawan/pkg/endeus"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg endeus.Config) (*gorm.DB, error) {
	log.Info().Msg("connecting to db...")
	gormConfig := &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}

	dsn := cfg.DBAddress()
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}
