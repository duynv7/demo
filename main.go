package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var cfg *Config
var messenger *Messenger

func init() {

	var err error
	cfg, err = loadConfigFromEnv()
	if err != nil {
		logrus.Fatalf("failed to load config %s\n", err)
	}
	messenger = &Messenger{
		AppSecret:   cfg.App.AppSecret,
		VerifyToken: cfg.App.AccessToken,
		Debug:       DebugAll,
	}

}

func main() {
	http.HandleFunc("/webhook", messenger.Handler)
	logrus.Infof("Messenger listens and serves at port %d\n", cfg.App.Port)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), nil))
}
