package config

import (
	"log"

	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/contracts/domain"

	"github.com/spf13/viper"
)

//LoadConfig loads the configuration for application
func LoadConfig(path string) domain.Config {
	viper.AddConfigPath(path)
	viper.SetConfigName(constants.ConfigName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	config := domain.Config{
		Server: domain.Server{
			Port: viper.Get("server.port").(string),
		},
		Database: domain.Database{
			Username: viper.Get("database.username").(string),
			Password: viper.Get("database.password").(string),
			Port:     viper.Get("database.port").(string),
			ConnStr:  viper.Get("database.connectionString").(string),
		},
		File: domain.File{
			Path: viper.Get("file.filePath").(string),
		},
	}
	return config
}
