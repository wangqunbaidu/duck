package server

import (
	"duck/internal/models"
	"duck/internal/pkg/config"
	"duck/internal/pkg/iplocator"
	"duck/internal/pkg/search"
	"duck/internal/scheduler"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	m "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/mlogclub/simple/common/strs"
	"github.com/mlogclub/simple/sqls"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init() {
	initConfig()
	initLogger()
	initDB()
	initCron()
	initIpLocator()
	initSearch()
}

func initConfig() {
	env := os.Getenv("BBSGO_ENV")
	if strs.IsBlank(env) {
		env = "dev"
	}

	viper.SetConfigName("duck." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.duck")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BBSGO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&config.Instance); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config.Instance.Env = env

	slog.Info("Load config", slog.String("ENV", env))
}

func initDB() {
	conf := config.Instance.DB
	db, err := gorm.Open(mysql.Open(conf.Url), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})

	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	if sqlDB, err := db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		sqlDB.SetConnMaxIdleTime(time.Duration(conf.ConnMaxIdleTimeSeconds) * time.Second)
		sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetimeSeconds) * time.Second)
	}

	// migrate
	if err := db.AutoMigrate(models.Models...); nil != err {
		slog.Error("auto migrate tables failed", slog.Any("error", err))
		panic(err)
	}
	if err := runMigrations(db); err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	sqls.SetDB(db)
}

func runMigrations(db *gorm.DB) error {
	s, _ := db.DB()
	driver, err := m.WithInstance(s, &m.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func initCron() {
	if config.Instance.IsProd() {
		scheduler.Start()
	}
}

func initIpLocator() {
	iplocator.InitIpLocator(config.Instance.IpDataPath)
}

func initSearch() {
	search.Init(config.Instance.Search.IndexPath)
}
