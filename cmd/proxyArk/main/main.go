package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8751")
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		log.Println(client.LocalAddr().String())
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	h, _ := url.ParseRequestURI(host)
	log.Println("adr:", hostPortURL.String(), ">>>", h)
	if hostPortURL.Opaque == "443" { // https access
		address = hostPortURL.String() // + ":443"
	} else { // http access
		if strings.Index(hostPortURL.Host, ":") == -1 { // host without port, default 80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}
	//After obtaining the requested host and port, start dialing
	log.Println("address", address)
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n")
	} else {
		server.Write(b[:n])
	} // Forwarding
	go io.Copy(server, client)
	io.Copy(client, server)
}
