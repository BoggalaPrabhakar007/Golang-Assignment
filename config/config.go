package config

import (
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/contracts/domain"
	"github.com/spf13/viper"
	"log"
)

//LoadConfig loads the configuration for application
func LoadConfig() domain.Config {
	viper.AddConfigPath(constants.ConfigPath)
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
		FilePath: domain.FilePath{
			FilePath: viper.Get("file.filePath").(string),
		},
	}
	return config
}
