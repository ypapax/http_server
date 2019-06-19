package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	addrToListen := "0.0.0.0:8080"
	glog.Infof("listening %+v", addrToListen)
	ln, err := net.Listen("tcp", addrToListen)
	if err != nil {
		glog.Exit(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			glog.Error(err)
			continue
		}
		go func() {
			if err := handleRequest(conn); err != nil {
				glog.Error(err)
			}
		}()

	}

}

func handleRequest(conn net.Conn) error {
	var buf = make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		glog.Error(err)
		return err
	}
	defer conn.Close()
	glog.V(4).Infof("read bytes: %+v, %+v", n, string(buf))

	if _, err := conn.Write([]byte("successfully read")); err != nil {
		glog.Error(err)
		return err
	}
	return nil
}
