package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"golang.org/x/net/proxy"
)

func copyAndLog(i net.Conn, o net.Conn) {
	written, _ := io.Copy(i, o)
	i.Close()
	o.Close()
	if verbose {
		fmt.Printf("Proxied %d bytes from %s to %s\n", written, i.RemoteAddr(), o.RemoteAddr())
	}
}

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

	if verbose {
		fmt.Printf("Client '%v' connected. Will proxy to %s\n", client.RemoteAddr(), targetHost)
	}

	target, err := dialSocksProxy.Dial("tcp", targetHost)
	if err != nil {
		log.Fatal("could not connect to target", err)
	}
	if verbose {
		fmt.Printf("Connection to server %v established\n", target.RemoteAddr())
	}

	go copyAndLog(target, client)
	go copyAndLog(client, target)
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
			log.Fatal("Could not accept client connection", err)
		}
		go handleConnection(proxyUser, proxyPassword, proxyHost, proxyPort, client, targetHost)
	}
}
