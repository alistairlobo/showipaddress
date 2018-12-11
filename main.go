package main

import (
    "net"
    "fmt"
    "log"
    "flag"
)


func main() {
    boolPtr := flag.Bool("help", false, "single binary ip address show tool, made by Egidijus. https://github.com/egidijus/showipaddress")
    flag.Parse()
    fmt.Println("help:", *boolPtr)
    ifaces, err := net.Interfaces()
    if err !=nil {
        log.Fatal(err)
    }
    // handle err
        for _, i := range ifaces {
            addrs, err := i.Addrs()
            if err !=nil {
                log.Fatal(err)
            }
            // handle err
            for _, addr := range addrs {
                var ip net.IP
                switch v := addr.(type) {
                case *net.IPNet:
                        ip = v.IP
                        fmt.Println(ip)
                case *net.IPAddr:
                        ip = v.IP
                        fmt.Println(ip)
                }
                // process IP address
            }
        }

}