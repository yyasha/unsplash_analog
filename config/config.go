package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWT_SECRET            string `mapstructure:"JWT_SECRET"`
	DB_ADDR               string `mapstructure:"DB_ADDR"`
	DB_NAME               string `mapstructure:"DB_NAME"`
	DB_USER               string `mapstructure:"DB_USER"`
	DB_PASSWORD           string `mapstructure:"DB_PASSWORD"`
	DB_MIGRATE_VERSION    uint   `mapstructure:"DB_MIGRATE_VERSION"`
	DOWN_OLD_DB_EVERYTIME bool   `mapstructure:"DOWN_OLD_DB_EVERYTIME"`
	RDB_PASSWORD          string `mapstructure:"RDB_PASSWORD"`
	RDB_ADDR              string `mapstructure:"RDB_ADDR"`
}

var envList []string = []string{
	"JWT_SECRET",
	"DB_ADDR",
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_MIGRATE_VERSION",
	"DOWN_OLD_DB_EVERYTIME",
	"RDB_PASSWORD",
	"RDB_ADDR",
}

// Global config variable
var Conf Config

// Load app config
func LoadConfig() (err error) {
	// viper.AutomaticEnv()

	for e := range envList {
		err = viper.BindEnv(envList[e])
		if err != nil {
			return err
		}
	}

	err = viper.Unmarshal(&Conf)
	return err
}
