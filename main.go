package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"
)

type PortResult struct {
	Host string `json:"host"`
	Ports []PortStatus `json:"ports"`
}

type PortStatus struct {
	Port string `json:"port"`
	Status string `json:"status"`
	Open bool `json:"open"`
}

func raw_connect(host string, port string) PortStatus {
	ps := PortStatus{}
	ps.Port = port

	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
			ps.Status = fmt.Sprint("CLOSED : ", err)
			ps.Open = false
	}
	if conn != nil {
			defer conn.Close()
			ps.Status = "OPEN"
			ps.Open = true
	}
	return ps
}

func main() {
	var(
		host = flag.String("h", "localhost", "target host")
		st = flag.Int("p", 58000, "start port")
		num = flag.Int("n", 1, "port nums")
		out_json = flag.Bool("json", false, "output json")
	)	
	flag.Parse()

	pr := PortResult{}
	pr.Host = *host

	if *num < 1 {
		if *out_json {
			fmt.Println("{}")
		}else{
			fmt.Println("num must > 0")
		}
		return
	}

	for i := *st; i < *st + *num; i++ {
		port := fmt.Sprintf("%d", i)
		ps := raw_connect(*host, port)
		pr.Ports = append(pr.Ports, ps)
	}

	if *out_json {
		outputJson, _ := json.Marshal(&pr)
		fmt.Println(string(outputJson))
	} else {	
		fmt.Println("Host:", *host)
		for _, ps := range pr.Ports {
			fmt.Printf("%s %s\n",ps.Port, ps.Status)
		}
	}
}