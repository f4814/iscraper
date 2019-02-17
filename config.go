package main

import (
	"time"
)

type CoreConfig struct {
	LogLevel string
}

type ScraperConfig struct {
	ID       int
	Cooldown time.Duration
	Scrapers int
	RootUser string
}

type DatabaseConfig struct {
	URL []string
	Authentication map[string]string
	Database string
}

type InstagramConfig struct {
	Username string
	Password string
	Proxy string
	CookieFile string
}

func NewCoreConfig(input map[string]interface{}) CoreConfig {
	return CoreConfig{
		LogLevel: input["log_level"].(string),
	}
}

func NewDatabaseConfig(input map[string]interface{}) DatabaseConfig {
	conf := DatabaseConfig{
		Database: input["database"].(string),
	}

	for _, v := range input["url"].([]interface{}) {
		conf.URL = append(conf.URL, v.(string))
	}

	auth := input["authentication"].(map[string]interface{})

	if auth != nil {
		conf.Authentication = make(map[string]string)
		conf.Authentication["username"] = auth["username"].(string)
		conf.Authentication["password"] = auth["password"].(string)
	}

	return conf
}

func NewInstagramConfig(input map[string]interface{}) InstagramConfig {
	return InstagramConfig{
		Username: input["username"].(string),
		Password: input["password"].(string),
		Proxy: input["proxy"].(string),
		CookieFile: input["cookie_file"].(string),
	}
}

func NewScraperConfig(input map[string]interface{}) ScraperConfig {
	duration, _ := time.ParseDuration(input["cooldown"].(string))
	return ScraperConfig{
		ID: 0,
		Cooldown: duration,
		Scrapers: input["scrapers"].(int),
		RootUser: input["root_user"].(string),
	}
}
