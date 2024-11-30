package ping

import (
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

// Функция проверяет доступность хоста
func pingTest(host string) bool {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		panic(err)
	}

	pinger.Count = 1
	pinger.Timeout = time.Millisecond * 300

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
