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
	h = flag.String("h","127.0.0.1:80","Target <host:port>.")
)

func init() {
	Cfg = new(config)
	flag.Parse()
	Cfg.Port = *p
	Cfg.HostIP = *h
}
