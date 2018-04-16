package config

import (
	"flag"
)

type config struct {
	Port string
	HostIP string
}

var (
	Cfg *config
	p = flag.String("p","8080","Listen on ?")
	h = flag.String("h","192.168.0.200:7001","Target host and port.")
)

func init() {
	Cfg = new(config)
	flag.Parse()
	Cfg.Port = *p
	Cfg.HostIP = *h
}
