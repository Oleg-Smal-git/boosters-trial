package main

import (
	"flag"
	"github.com/Oleg-Smal-git/boosters-trial/app/config"
	"strconv"

	log "github.com/sirupsen/logrus"
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
		p, err := strconv.Atoi(cfg["self.port"])
		if err != nil {
			panic(err)
		}
		port = p
	}
	log.Infof("starting the app at %v:%v", host, port)
}
