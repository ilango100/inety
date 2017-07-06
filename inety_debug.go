// +build debug

package main

import ( "fmt"
	"net"
	"io"
	"os"
	"sync"
)

func main() {
	var addr,port string
	var wg sync.WaitGroup

	fmt.Println("Enter addr,port")
	fmt.Scan(&addr,&port)
	fmt.Println("Your input addr: ",addr," port: ",port)

	conn, err := net.Dial("tcp",net.JoinHostPort(addr,port))
	defer conn.Close()

	if err != nil {
		fmt.Println("Error while connecting...",err)
	}
	fmt.Println("connected")

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _,err = io.Copy(os.Stdout,conn) ; err == nil; _,err = io.Copy(os.Stdout,conn) {
			//nothing to do
		}
		fmt.Println("Output is over :",err)
	}()

	go func() {
		defer wg.Done()
		for _,err = io.Copy(conn,os.Stdin) ; err ==nil; _,err = io.Copy(conn,os.Stdin) {
			//nothing to do actually
		}
		fmt.Println("Input is over: ",err)
	}()

	if err != nil { fmt.Println("Error: ",err) }

	wg.Wait()
}
