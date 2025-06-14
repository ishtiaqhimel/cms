package db

import (
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	"github.com/ishtiaqhimel/news-portal/cms/internal/config"
)

type PgClient struct {
	*gorm.DB
}

// pg is the postgres instance
var pg = &PgClient{}

// Get returns the default pgClient currently in use
func Get() *PgClient {
	return pg
}

// Connect database, must call once before server boot to Get() the pg instance
func Connect() (err error) {
	if pg.DB != nil {
		logrus.Info("postgres already initialized")
		return nil
	}
	cfg := config.Get().Database
	uriFormat := fmt.Sprintf("host=%%s port=%%d sslmode=%s user=%s password=%s dbname=%s",
		cfg.SslMode,
		cfg.Username,
		cfg.Password,
		cfg.Name,
	)
	primaryURI := fmt.Sprintf(uriFormat, cfg.Primary.Host, cfg.Primary.Port)
	secondaryURI := fmt.Sprintf(uriFormat, cfg.Secondary.Host, cfg.Secondary.Port)

	logMode := logger.Info
	if cfg.Debug {
		logMode = logger.Info
	}

	pg.DB, err = gorm.Open(postgres.Open(primaryURI), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return fmt.Errorf("failed to open gorm conn: %v", err)
	}

	sqlDB, err := pg.DB.DB()
	if err != nil {
		return err
	}

	// connection pool settings
	resolver := dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{postgres.Open(secondaryURI)},
	})
	if cfg.MaxIdleConn != 0 {
		resolver.SetMaxIdleConns(cfg.MaxIdleConn)
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)
	}
	if cfg.MaxOpenConn != 0 {
		resolver.SetMaxOpenConns(cfg.MaxOpenConn)
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	}
	if cfg.MaxLifeTime.Seconds() != 0 {
		resolver.SetConnMaxLifetime(cfg.MaxLifeTime * time.Second)
		sqlDB.SetConnMaxLifetime(cfg.MaxLifeTime * time.Second)
	}

	return pg.Use(resolver)
}

func Close() error {
	sqlDB, err := pg.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func Migrate(uri, path string) error {
	if uri == "" {
		cfg := config.Get().Database
		uri = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Primary.Host,
			cfg.Primary.Port, cfg.Name, cfg.SslMode)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s", path),
		uri)
	if err != nil {
		return err
	}
	return m.Up()
}
