package statistic

import (
	"context"
	probing "github.com/prometheus-community/pro-bing"
	"time"
)

func pingThread(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Инициализируем пингер
			pinger, err := probing.NewPinger("8.8.8.8")
			if err != nil {
				panic(err)
			}

			// Устанавливаем количество попыток
			pinger.Count = 4

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

			// Сохраняем задержку в переменную
			StatisticMutex.Lock()
			if stats.PacketsRecv != 0 {
				PingStatus = uint64(stats.MaxRtt.Milliseconds())
			} else {
				PingStatus = 0
			}
			StatisticMutex.Unlock()

			// Ждём до следующей итерации
			time.Sleep(time.Millisecond * 100)
		}
	}
}
