package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/gen2brain/beeep"
)

// getLocalIP retrieves the local IP address by establishing a UDP connection
func getLocalIP() (string, error) {
	// Connect to an external address to determine local IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func playBeep() {
	beeep.Beep(44100, 400)
}

func main() {

	var beepPntr *bool = flag.Bool("beep", false, "Play beep")
	flag.Parse()

	var previousIP string

	currentIP, err := getLocalIP()
	if err != nil {
		fmt.Println("Error getting local IP:", err)
		time.Sleep(10 * time.Second)
	}

	previousIP = currentIP

	fmt.Println("\033[032m" + currentIP + "\033[0m")
	if *beepPntr {
		playBeep()
	}
	for {
		currentIP, err := getLocalIP()
		if err != nil {
			fmt.Println("Error getting local IP:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		if currentIP != previousIP {
			fmt.Printf("\033[031m  ->: %s\n", currentIP)
			previousIP = currentIP
			if *beepPntr {
				playBeep()
			}
		}
		time.Sleep(10 * time.Second)
	}
}
