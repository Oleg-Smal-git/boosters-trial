package main

import (
	"flag"
	"github.com/Oleg-Smal-git/boosters-trial/app/config"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func main() {
	// Load configs.
	cfg := config.MustConfig()
	log.Infof("starting the app with %v env", cfg["envname"])

	// Parse override flags.
	var host string
	var port int
	flag.StringVar(&host, "host", "", "override host, defaulted to config if not provided")
	flag.IntVar(&port, "port", 0, "override port, defaulted to config if not provided")
	flag.Parse()
	if host == "" {
		host = cfg["self.host"]
	}
	if port == 0 {
		port, _ = strconv.Atoi(cfg["self.port"])
	}
	log.Infof("starting the app at %v:%v", host, port)
}
