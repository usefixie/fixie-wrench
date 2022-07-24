package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseConnectionString(connectionString string) (string, string, string, int) {
	re := regexp.MustCompile(`[:@]`)
	split := re.Split(connectionString, -1)
	if len(split) != 4 {
		println("Fixie Socks connection string was provided but invalid. Expected a connection string in the form of fixie:TOKEN@FIXIE_HOST:FIXIE_PORT. This can be provided in the FIXIE_SOCKS_HOST environment variable or fixieSocksHost command line flag.")
		os.Exit(1)
	}
	proxyPort, err := strconv.Atoi(split[3])
	if err != nil {
		println("Fixie Socks proxy port must be an interger")
		os.Exit(1)
	}
	return split[0], split[1], split[2], proxyPort
}

func getSocksConnection(socksConnectionString string) (string, string, string, int) {
	if socksConnectionString == "" {
		socksConnectionString = os.Getenv("FIXIE_SOCKS_HOST")
	}
	if socksConnectionString == "" {
		fmt.Println("One of the FIXIE_SOCKS_HOST environment variable or fixieSocksHost command line flag must be set")
		os.Exit(1)
	}
	return parseConnectionString(socksConnectionString)
}
