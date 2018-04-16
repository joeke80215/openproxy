package core

import (
	"net"
	"log"
	"io"
)

func copy(wc io.WriteCloser, r io.Reader) {
	defer wc.Close()
	io.Copy(wc, r)
}

func handleClinet (us net.Conn, server string) {
	ds, err := net.Dial("tcp", server)
	if err != nil {
		us.Close()
		log.Printf("failed to dial %s: %s", server, err)
		return
	}

	go copy(ds, us)
	go copy(us, ds)
}
