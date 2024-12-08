package scan

import (
	"net"
	"strconv"
	"sync"
	"time"
)

// Фукнция сканирует порты конкретного адреса
func scanHost(ip string, PortsCount *int, mutex *sync.Mutex) []int {
	openPorts := make([]int, 0)

	for port := 1; port <= 1024; port++ {
		d := net.Dialer{Timeout: time.Millisecond * 300}
		conn, err := d.Dial("tcp", ip+":"+strconv.Itoa(port))

		mutex.Lock()
		*PortsCount += 1
		mutex.Unlock()
		if err == nil {
			conn.Close()
			openPorts = append(openPorts, port)
		}
	}

	return openPorts
}
