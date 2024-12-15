package ping

import (
	"QuickProbe/pkg/argflags"
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

// Функция проверяет доступность хоста
func pingTest(host string) bool {
	// Инициализируем пингер
	pinger, err := probing.NewPinger(host)
	if err != nil {
		panic(err)
	}

	// Устанавливаем тайм аут для пинга
	pinger.Timeout = time.Millisecond * time.Duration(*argflags.Timeout)

	// Устанавливаем количество попыток
	pinger.Count = 1

	// Устанавливаем непревилигированный доступ
	pinger.SetPrivileged(false)

	// Отключение отладочной информации
	pinger.Debug = false

	// Запускаем пингование
	err = pinger.Run()
	if err != nil {
		panic(err)
	}

	// Получаем статистику
	stats := pinger.Statistics()

	// Возвращаем результат
	if stats.PacketsRecv != 0 {
		return true
	} else {
		return false
	}
}
