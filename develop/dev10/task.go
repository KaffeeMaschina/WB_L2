package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
type Config struct {
	host    string
	port    string
	timeout time.Duration
}

func NewConfig() *Config {
	timeout := flag.Duration("timeout", time.Second*10, "timeout for connection")

	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}
	host := flag.Arg(0)
	port := flag.Arg(1)
	return &Config{
		host:    host,
		port:    port,
		timeout: *timeout,
	}
}
func main() {
	config := NewConfig()
	fmt.Println(config)
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(config.host, config.port), config.timeout)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		errClose := conn.Close()
		if errClose != nil {
			log.Fatalln("error closing connection: ", err.Error())
		}
	}(conn)

	fmt.Println("connected to server")
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go Read(conn)
	go Write(conn)

	<-sigCh
	fmt.Println("Closing connection...")

}
func Write(conn net.Conn) {
	for {
		var b [1]byte
		_, err := os.Stdin.Read(b[:])
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from stdin:", err)
			os.Exit(1)
		}

		// If Ctrl+D is pressed, close the connection
		if b[0] == 4 {
			fmt.Println("Ctrl+D pressed. Closing connection...")
			conn.Close()
			return
		}

		// Write the byte to the connection
		_, err = conn.Write(b[:])
		if err != nil {
			fmt.Println("Error writing to server:", err)
			return
		}
	}

}
func Read(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading from server:", err)
			os.Exit(1)
		}
		fmt.Print(string(buf[:n]))

	}
}
