package main

import (
	probing "github.com/prometheus-community/pro-bing"
	"net"
	"strconv"
	"time"
)

// ScannerThread является сканирующей горутиной, сюда попадают адреса после пингования
func ScannerThread(IPChannel chan string) {
	defer WorkWG.Done()

	for {
		ip, ok := <-IPChannel
		if ok {
			ports := scanHost(ip)
			if len(ports) != 0 {
				for _, port := range ports {
					println(ip, port)
				}
			}
		} else {
			break
		}
	}
}

// Фукнция сканирует порты конкретного адреса
func scanHost(ip string) []int {
	openPorts := make([]int, 0)

	for port := 1; port <= 1024; port++ {
		d := net.Dialer{Timeout: time.Millisecond * 100}
		conn, err := d.Dial("tcp", ip+":"+strconv.Itoa(port))

		if err == nil {
			conn.Close()
			openPorts = append(openPorts, port)
		}
	}

	return openPorts
}

// PingingThread является пингующий горутиной, сначала сюда попадают адреса для проверки
func PingingThread(RawIPChannel chan string, IPChannel chan string) {
	defer PingWG.Done()

	for {
		ip, ok := <-RawIPChannel
		if ok {
			available := pingTest(ip)
			if available {
				IPChannel <- ip
			}
		} else {
			break
		}
	}
}

// Функция проверяет доступность хоста
func pingTest(host string) bool {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		panic(err)
	}

	pinger.Count = 1
	pinger.Timeout = time.Millisecond * 100

	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()
	if stats.PacketsRecv != 0 {
		return true
	} else {
		return false
	}
}
