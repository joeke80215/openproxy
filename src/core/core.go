package core

import (
	"net"
	"fmt"
	"log"
	"../cmd"
)

func StartProxy () {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%s", cmd.Cfg.Port))
	if err != nil {
		log.Fatalf("failed to bind: %s", err)
	}

	log.Printf("listening on %s, proxy %s", cmd.Cfg.Port, cmd.Cfg.HostIP)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("failed to accept: %s", err)
			continue
		}
		log.Println("Remote IP : ",conn.RemoteAddr().String(),)
		go handleClinet(conn, cmd.Cfg.HostIP)
	}
}
