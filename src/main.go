package main

import (
	"QuickProbe/pkg/argflags"
	"QuickProbe/pkg/network"
	"QuickProbe/pkg/ping"
	"QuickProbe/pkg/scan"
	"QuickProbe/pkg/statistic"
	"context"
)

// Создаём указатели на общие обьекты
var (
	IPChannel    chan string // Канал для передачи проверенных адресов
	RawIPChannel chan string // Канал для передачи непроверенных адресов
)

func main() {
	// Получаем флаги
	argflags.InitFlags()
	argflags.ParseFlags()
	argflags.CheckFlags()

	// Создаём сеть
	net := network.NewNetwork(*argflags.InputNet)

	// Подключаем базу данных

	// Инициализируем канал для передачи проверенных адресов
	IPChannel = make(chan string, *argflags.AddressBufferSize)

	// Инициализирую канал для передачи непроверенных адресов
	RawIPChannel = make(chan string, *argflags.AddressBufferSize)

	// Инициализируем контексты для подсистем
	statCtx, statCancel := context.WithCancel(context.Background())

	// Инициализируем поток статистики
	statistic.StatisticStart(statCtx, IPChannel, RawIPChannel)

	// Инициализируем рабочие потоки
	for i := uint64(0); i < *argflags.NumberScanThreads; i++ {
		scan.WorkWG.Add(1)
		go scan.ScannerThread(IPChannel)
	}

	// Инициализируем пингующие потоки
	for i := uint64(0); i < *argflags.NumberPingThreads; i++ {
		ping.PingWG.Add(1)
		go ping.PingingThread(RawIPChannel, IPChannel)
	}

	// Запускаем основную петлю для генерации и передачи адресов
	for !net.Ended {
		if !net.IsPrivate() {
			RawIPChannel <- net.String()
		}

		net.Inc()
	}

	// Закрываем канал пингующих потоков
	for {
		if len(RawIPChannel) == 0 {
			close(RawIPChannel)
			break
		}
	}

	// Ожидаем завершение работы пингующих потоков
	ping.PingWG.Wait()

	// Закрываем канал сканирующих потоков
	for {
		if len(IPChannel) == 0 {
			close(IPChannel)
			break
		}
	}

	// Ожидаем завершение работы сканирующих потоков
	scan.WorkWG.Wait()
	statCancel()
}
