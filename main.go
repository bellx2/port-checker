package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func raw_connect(host string, ports []string) {
	for _, port := range ports {
			timeout := time.Second
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
			if err != nil {
					fmt.Println(port, ": CLOSE - ", err)
			}
			if conn != nil {
					defer conn.Close()
					fmt.Println(port, ": OPEN")
			}
	}
}

func main() {
	var(
		host = flag.String("h", "localhost", "target host")
		st = flag.Int("p", 58000, "start port")
		num = flag.Int("n", 1, "port nums")
	)	
	flag.Parse()
	
	if *num < 1 {
		fmt.Println("num must > 0")
		return
	}

	fmt.Println("Host:", *host)
	for i := *st; i < *st + *num; i++ {
		port := fmt.Sprintf("%d", i)
		raw_connect(*host, []string{port})
	}
}