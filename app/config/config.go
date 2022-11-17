package config

import (
	"github.com/Oleg-Smal-git/boosters-trial/app/helpers"
	"os"
	"strings"

	"github.com/robfig/config"
)

var (
	// envCache stores env flag value.
	envCache string
	// configCache stores the value of LoadConfig after its first call.
	configCache map[string]string
)

// Config fetches project configs based on current ENV, caching the result.
func Config() (map[string]string, error) {
	if configCache == nil {
		cfg, err := evaluateConfig(env())
		if err != nil {
			return nil, err
		}
		configCache = cfg
	}
	return configCache, nil
}

// evaluateConfig loads configs based on ENV flag.
func evaluateConfig(env string) (map[string]string, error) {
	cfg, err := config.ReadDefault(helpers.BasePath() + "/config/app.conf")
	if err != nil {
		return nil, err
	}
	options, err := cfg.Options(env)
	if err != nil {
		return nil, err
	}
	params := make(map[string]string)
	for _, o := range options {
		value, err := cfg.String(env, o)
		if err != nil {
			return nil, err
		}
		params[o] = strings.Trim(value, "\"")
	}
	return params, nil
}

// env fetches current env.
func env() string {
	if envCache == "" {
		envCache = evaluateEnv()
	}
	return envCache
}

// evaluateEnv fetches ENV flag from environment variables. Returns "default" if one is fails to be found.
func evaluateEnv() string {
	if env := os.Getenv("ENV"); env != "" {
		return env
	}
	return "default"
}
