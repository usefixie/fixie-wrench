package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"golang.org/x/net/proxy"
)

func handleConnection(proxyUser string, proxyPassword string, proxyHostName string, proxyPort int, client net.Conn, targetHost string) {
	baseDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 60 * time.Second,
	}
	proxyHost := fmt.Sprintf("%s:%d", proxyHostName, proxyPort)
	proxyAuth := proxy.Auth{User: proxyUser, Password: proxyPassword}
	dialSocksProxy, err := proxy.SOCKS5("tcp", proxyHost, &proxyAuth, baseDialer)
	if err != nil {
		panic("Error creating SOCKS proxy")
	}

	fmt.Printf("client '%v' connected!\n", client.RemoteAddr())

	target, err := dialSocksProxy.Dial("tcp", targetHost)
	if err != nil {
		log.Fatal("could not connect to target", err)
	}
	fmt.Printf("connection to server %v established!\n", target.RemoteAddr())

	go func() { io.Copy(target, client) }()
	go func() { io.Copy(client, target) }()
}

func startServer(proxyUser string, proxyPassword string, proxyHost string, proxyPort int, localPort int, targetHost string) {
	wg.Add(1)
	incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", localPort))
	if err != nil {
		log.Fatalf("Could not listen on %d: %v", localPort, err)
	}
	defer wg.Done()

	for {
		client, err := incoming.Accept()
		if err != nil {
			log.Fatal("could not accept client connection", err)
		}
		go handleConnection(proxyUser, proxyPassword, proxyHost, proxyPort, client, targetHost)
	}
}
