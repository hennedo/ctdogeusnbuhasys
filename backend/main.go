package main

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/henne-/ctdogeusnbuhasys/backend/db"
	_ "github.com/henne-/ctdogeusnbuhasys/backend/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var waitGroup sync.WaitGroup

func main() {
	db.Load()

	initConfig()
	initServer()

	loadConfig()
	startServer(&waitGroup)
	waitGroup.Wait()
	db.Close()
}

func initConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("Config file not found. Using defaults!")
		} else {
			logrus.Fatalf("Fatal error config file: %s", err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("Config file changed: %s", e.Name)
		loadConfig()
		startServer(&waitGroup)
	})

	viper.AutomaticEnv()

	viper.SetDefault("http.ListenAddress", "127.0.0.1:8080")
}

func loadConfig() {
	httpListenAddress = viper.GetString("http.ListenAddress")
}
