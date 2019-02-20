package github.com/hantmac/juejinxiaoceToMobi/config

import (
	"encoding/json"
	"os"
	"sync/atomic"
)

var cfgValue atomic.Value

type Config struct {
	cfgFile  string
	GetSectionUrl string
	GetUrl string
	Token    string `json:"token"`
	ClientID string `json:"client_id"`
	Src      string `json:"src"`
	UID      string `json:"uid"`
	ID       string `json:"id"`
	Title string
}

func Cfg() *Config {
	return cfgValue.Load().(*Config)
}

func init() {

	cfg := &Config{
		cfgFile: "./config/config.json",
	}
	if err := cfg.load(); err != nil {
		panic(err)
	}
}

func (cfg *Config) load() error {

	file, err := os.Open(cfg.cfgFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(cfg); err != nil {
		panic(err)
	}

	cfgValue.Store(cfg)

	return nil
}
