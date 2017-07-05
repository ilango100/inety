package main

import ( "fmt"
	"net"
	"io"
	"os"
	"bufio"
)

func main() {
	var addr,port string
	fmt.Println("Enter addr,port")
	fmt.Scan(&addr,&port)
	fmt.Println("Your input addr: ",addr," port: ",port)

	conn, err := net.Dial("tcp",net.JoinHostPort(addr,port))
	defer conn.Close()

	if err != nil {
		fmt.Println("Error while connecting...")
	}
	fmt.Println("connected")

	notif := make(chan bool)

	go func() {
		for _,err = io.Copy(os.Stdout,conn) ; err == nil; {
			_,err = io.Copy(os.Stdout,conn)
		}
		fmt.Println("Output is over :",err)
		notif <- true
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		writer := bufio.NewWriter(conn)
		for s,err := reader.ReadString('\n') ; err == nil; s,err = reader.ReadString('\n') {
			writer.WriteString(s)
			writer.Flush()
		}
		fmt.Println("Input is over: ",err)
		notif <- true
	}()


	if err != nil { fmt.Println("Error: ",err) }

	<-notif
	<-notif

}
