package main

import (
	"flag"
	"strconv"

	"github.com/Oleg-Smal-git/boosters-trial/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Load configs.
	cfg, err := config.Config()
	if err != nil {
		panic(err)
	}
	log.Infof("starting the app with %v env", cfg["envname"])

	// Parse override flags.
	var host string
	var port int
	flag.StringVar(&host, "host", "", "override host, defaulted to config if not provided")
	flag.IntVar(&port, "port", 0, "override port, defaulted to config if not provided")
	flag.Parse()
	if host == "" {
		host = cfg["host"]
	}
	if port == 0 {
		p, err := strconv.Atoi(cfg["port"])
		if err != nil {
			panic(err)
		}
		port = p
	}
	log.Infof("starting the app at %v:%v", host, port)
}
