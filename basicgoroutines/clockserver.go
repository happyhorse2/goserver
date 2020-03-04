package basicgoroutines

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type person struct {
	age  int
	name string
}

// Clockserver : 串行clock服务器
func Clockserver() {
	fmt.Println("server start")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	var personone person
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		//go  handleConn(conn) // 并发处理多个连接
		go handleConnv2(conn, &personone) //处理一个连接，处理方式有阻塞
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func handleConnv2(c net.Conn, personptr *person) {
	defer c.Close()
	personptr.age = personptr.age + 1 //非局部变量，多个协程共同使用
	fmt.Println(personptr.age)
	input := bufio.NewScanner(c)
	for input.Scan() {
		string := input.Text()
		fmt.Println(string)
		go echo(c, string, 10*time.Second) //处理一个连接阻塞
	}
	fmt.Println("server end")
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	fmt.Println(strings.ToUpper(shout))
	c.Close()
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	fmt.Println(shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Println(strings.ToLower(shout))
}

func testFprintln() {

}
