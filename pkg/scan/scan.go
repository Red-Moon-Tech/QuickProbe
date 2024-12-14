package scan

import (
	"QuickProbe/pkg/argflags"
	"QuickProbe/pkg/statistic"
	"net"
	"strconv"
	"time"
)

// Фукнция сканирует порты конкретного адреса
func scanHost(ip string) []int {
	openPorts := make([]int, 0)

	for port := 1; port <= 1024; port++ {
		d := net.Dialer{Timeout: time.Millisecond * time.Duration(*argflags.Timeout)}
		conn, err := d.Dial("tcp", ip+":"+strconv.Itoa(port))

		statistic.StatisticMutex.Lock()
		statistic.PortsCounter += 1
		statistic.StatisticMutex.Unlock()
		if err == nil {
			conn.Close()
			openPorts = append(openPorts, port)
		}
	}

	return openPorts
}
