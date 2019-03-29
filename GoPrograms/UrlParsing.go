package main

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func main() {

	s := "http://localhost:8080/Go/api/version2/divyaname/info?status=123456789"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("u.Scheme   ", u.Scheme)
	fmt.Println("u.User  ", u.User)

	fmt.Println("u.Host  ", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host  ", host)
	fmt.Println("port  ", port)
	fmt.Println("u.Path  ", u.Path)
	fmt.Println("u.Fragment  ", u.Fragment)
	fmt.Println("u.RawQuery  ", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("m    ", m)
	fmt.Println("status  ", m["status"][0])

	//from /Go/api/version2/divyaname/info ,getting only divyaname
	url := "/Go/api/version2/divyaname/info"
	ss := strings.TrimSuffix(url, "/info")
	t := strings.TrimPrefix(ss, "/Go/api/version2/")
	//fmt.Println(s)
	fmt.Println(t)

}
