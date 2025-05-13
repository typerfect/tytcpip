package main

import (
	"log"
	"os/exec"
	"tytcpip/tcpip/link/tap"
)

func main() {
	tap, err := tap.NewTap("tap0")
	if err != nil {
		log.Fatalln("new tap failed, ", err)
	}

	if err := tap.Linkup(); err != nil {
		log.Fatalln("tap linkup failed", err)
	}

	if err := tap.SetIp("192.168.3.105"); err != nil {
		log.Fatalln("set ip failed", err)
	}

	out, err := exec.Command("ifconfig").Output()
	if err != nil {
		log.Fatalln("ifconfig failed,", out)
	}
	log.Println("ifconfig", string(out))
}
