package scan

import (
	"QuickProbe/pkg/argflags"
	"QuickProbe/pkg/results"
	"QuickProbe/pkg/statistic"
	"net"
	"strconv"
	"time"
)

// Фукнция сканирует порты конкретного адреса
func scanHost(ip string) {
	// Открытые порты найденные в результате сканирования
	openPorts := make([]int, 0)

	// Получаем список портов
	var portsArray = make([]int, 0)
	portsArray = append(portsArray, argflags.PortsList...)

	for _, port := range portsArray {
		// Подключаемся к хосту
		d := net.Dialer{Timeout: time.Millisecond * time.Duration(*argflags.Timeout)}
		conn, err := d.Dial("tcp", ip+":"+strconv.Itoa(port))

		// Увеличиваем счётчик отсканированных портов
		statistic.StatisticMutex.Lock()
		statistic.PortsCounter += 1
		statistic.StatisticMutex.Unlock()

		// Если соединение было успешно, то закрываем его и добавляем порт в список
		if err == nil {
			conn.Close()
			openPorts = append(openPorts, port)
		}
	}
	results.ResMutex.Lock()
	if len(openPorts) > 0 {
		results.ResultMap[ip] = openPorts
	}
	results.ResMutex.Unlock()
}
