package main

import (
	"QuickProbe/pkg/argflags"
	"QuickProbe/pkg/network"
	"QuickProbe/pkg/ping"
	"QuickProbe/pkg/results"
	"QuickProbe/pkg/scan"
	"QuickProbe/pkg/statistic"
	"context"
	"fmt"
	"os"
	"time"
)

// Создаём общие обьекты
var (
	ScanIPChannel chan string // Канал для передачи адресов для сканирования
	PingIPChannel chan string // Канал для передачи адресов для пингования
)

func main() {
	// Получаем флаги
	argflags.InitFlags()
	argflags.ParseFlags()
	argflags.CheckFlags()

	// Проверяем если флаг Network - файл
	if _, err := os.Stat(*argflags.InputNet); err == nil {
		data, _ := os.ReadFile(*argflags.InputNet)
		*argflags.InputNet = string(data)
	}

	// Создаём сеть
	net := network.NewNetwork(*argflags.InputNet)

	// Инициализируем пакет результатов
	results.Init()

	// Подключаем базу данных

	// Инициализируем канал для передачи адресов для сканирования
	ScanIPChannel = make(chan string, *argflags.AddressBufferSize)

	// Инициализирую канал для передачи адресов для пингования
	PingIPChannel = make(chan string, *argflags.AddressBufferSize)

	// Инициализируем контексты для подсистем
	statCtx, statCancel := context.WithCancel(context.Background())

	// Инициализируем поток статистики
	statistic.StatisticStart(statCtx, ScanIPChannel, PingIPChannel)

	// Инициализируем рабочие потоки
	if *argflags.SkipPingThreads {
		// Если флаг SkipPingThreads == true, то берем хосты прямиком из PingIpChannel
		for i := uint64(0); i < *argflags.NumberScanThreads; i++ {
			scan.WorkWG.Add(1)
			go scan.ScannerThread(PingIPChannel)
		}
	} else {
		for i := uint64(0); i < *argflags.NumberScanThreads; i++ {
			scan.WorkWG.Add(1)
			go scan.ScannerThread(ScanIPChannel)
		}
	}

	// Инициализируем пингующие потоки
	if !(*argflags.SkipPingThreads) {
		for i := uint64(0); i < *argflags.NumberPingThreads; i++ {
			ping.PingWG.Add(1)
			go ping.PingingThread(PingIPChannel, ScanIPChannel)
		}
	}

	// Запоминаем время начала сканирования
	timeStart := time.Now()

	// Запускаем основную петлю для генерации и передачи адресов
	for !net.Ended {
		PingIPChannel <- net.String()
		net.Inc()
	}

	// Закрываем канал пингующих потоков
	for {
		if len(PingIPChannel) == 0 {
			close(PingIPChannel)
			break
		}
	}

	// Ожидаем завершение работы пингующих потоков
	ping.PingWG.Wait()

	// Закрываем канал сканирующих потоков
	for {
		if len(ScanIPChannel) == 0 {
			close(ScanIPChannel)
			break
		}
	}

	// Ожидаем завершение работы сканирующих потоков
	scan.WorkWG.Wait()

	// Запоминаем время конца сканирования
	timeEnd := time.Now()

	// Завершаем потоки связанные с работой подсистемы сбора статистики
	statCancel()

	// Выводим результаты сканирования
	results.ShowResults()

	// Выводим время сканирования
	fmt.Printf("Scan time was %.2f seconds\n", timeEnd.Sub(timeStart).Seconds())

}
