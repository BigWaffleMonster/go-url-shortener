package main

import (
	"log"

	"gihtub.com/BigWaffleMonster/go-url-shortener"
  DB "gihtub.com/BigWaffleMonster/go-url-shortener/pkg/db"
	"gihtub.com/BigWaffleMonster/go-url-shortener/pkg/handler"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error with initializing configs: %s", err.Error())
	}

  db, err := DB.NewPostgresDB(DB.Config{
    Host: viper.GetString("db.host"),
    Port: viper.GetString("db.port"),
    Username: viper.GetString("db.username"),
    DBName: viper.GetString("db.dbname"),
    SSLMode: viper.GetString("db.sslmode"),
    Password: viper.GetString("db.password"),
  })

  if err != nil {
    log.Fatalf("Failed to initialize db: %s", err.Error())
  }

	srv := new(urlshortener.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes(db)); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
