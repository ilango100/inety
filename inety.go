// +build !debug

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func parseArgs() (addr, port string) {

	var help, h bool

	flag.BoolVar(&help, "help", false, "Show usage")
	flag.BoolVar(&h, "h", false, "Show usage")
	flag.StringVar(&port, "p", "", "Port number")
	flag.StringVar(&addr, "a", "", "Address to connect to")
	flag.Parse()

	if help || h {
		flag.Usage()
		os.Exit(0)
	}

	if (addr == "" || port == "") && flag.NArg() > 0 {
		switch flag.NArg() {
		default:
			fallthrough
		case 2:
			port = flag.Arg(1)
			fallthrough
		case 1:
			addr = flag.Arg(0)
		}
	}
	return

}

func main() {
	var addr, port string = parseArgs()
	var wg sync.WaitGroup

	if addr == "" || port == "" {
		fmt.Println("Enter addr,port")
		fmt.Scan(&addr, &port)
	}

	conn, err := net.Dial("tcp", net.JoinHostPort(addr, port))
	defer conn.Close()

	if err != nil {
		fmt.Println("Error while connecting...", err)
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		buf := make([]byte, 512, 512)
		for nr, err := conn.Read(buf); err == nil; nr, err = conn.Read(buf) {
			nw, err := os.Stdout.Write(buf[:nr])
			if nr != nw && err != nil {
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		buf := make([]byte, 512, 512)
		for nr, err := os.Stdin.Read(buf); err == nil; nr, err = os.Stdin.Read(buf) {
			nw, err := conn.Write(buf[:nr])
			if nr != nw && err != nil {
				break
			}
		}
	}()

	wg.Wait()
}
