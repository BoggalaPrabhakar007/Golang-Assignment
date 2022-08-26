package config

import (
	"log"
	"reflect"
	"testing"

	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/constants"
	"github.com/BoggalaPrabhakar007/golang-assignment/pkg/contracts/domain"

	"github.com/spf13/viper"
)

func TestLoadConfig(t *testing.T) {
	viper.AddConfigPath("../config")
	viper.SetConfigName(constants.ConfigName)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want domain.Config
	}{
		{
			name: "config test",
			args: args{path: "../config"},
			want: domain.Config{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
