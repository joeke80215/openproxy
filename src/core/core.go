package core

import (
	"net"
	"fmt"
	"log"
	"../config"
)

func StartProxy () {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s",config.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to bind: %s", err)
	}

	log.Printf("listening on %s, proxy %s", config.Cfg.Port, config.Cfg.HostIP)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept: %s", err)
			continue
		}
		log.Println("Remote IP : ",conn.RemoteAddr().String(),)
		go handleClinet(conn, config.Cfg.HostIP)
	}
}
