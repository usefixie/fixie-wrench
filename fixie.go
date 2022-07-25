package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
)

var verbose = false
var wg = &sync.WaitGroup{}

func parseForwardingArg(arg string) (int, string, int) {
	s := strings.Split(arg, ":")
	if len(s) != 3 {
		fmt.Print("Invalid port fowarding argument. Each positional port forwarding argument should be in the form of localPort:remoteHost:remotePort\n")
		os.Exit(1)
	}
	localPort, err := strconv.Atoi(s[0])
	if err != nil {
		fmt.Printf("Invalid local port '%s'. Port numbers should be integers.\n", s[0])
		os.Exit(1)
	}
	remoteHost := s[1]
	remotePort, err := strconv.Atoi(s[2])
	if err != nil {
		fmt.Printf("Invalid remote port '%s'. Port numbers should be integers.", s[0])
		os.Exit(1)
	}
	return localPort, remoteHost, remotePort
}

func main() {

	// Listen for keyboard interupt
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		for _ = range signals {
			fmt.Println("\nReceived an interrupt, stopping...")
			os.Exit(0)
		}
	}()

	// Parse command line arguments and Fixie Socks
	verboseFlag := flag.Bool("v", false, "Specifies verbose mode")
	var socksConnectionString = flag.String("fixieSocksHost", "", "[Optional] The Fixie Socks connection string. If not provided, will use eng.FIXIE_SOCKS_HOST")
	flag.Parse()
	proxyUser, proxyPassword, proxyHost, proxyPort := getSocksConnection(*socksConnectionString)
	verbose = *verboseFlag

	if verbose {
		fmt.Println("Fixie CLI (Verbose Mode)")
		fmt.Printf("Fixie Socks cluster: %s:%d\n", proxyHost, proxyPort)
	}

	// Listen for connections and forward requests via Fixie Socks proxy
	args := flag.Args()
	if flag.NArg() == 0 {
		fmt.Println("No positional arguments provided for Fixie Socks port forwarding. Expected one or more positional arguments in the form localPort:remoteHost:remotePort. For more information: https://usefixie.com/documentation/socks")
		os.Exit(1)
	}
	for _, arg := range args {
		wg.Add(1)
		localPort, remoteHost, remotePort := parseForwardingArg(arg)
		targetHost := fmt.Sprintf("%s:%d", remoteHost, remotePort)
		fmt.Printf("Forwarding local port %d to %s via Fixie Socks\n", localPort, targetHost)
		go startServer(proxyUser, proxyPassword, proxyHost, proxyPort, localPort, targetHost)
	}
	wg.Wait()
}
