package main

import (
	"log"
	"sync"
)

// Инициализируем переменные под хранение флагов
var (
	InputNet          *string // Сеть для сканирования
	NumberScanThreads *uint64 // Количество потокв для сканирования
	NumberPingThreads *uint64 // Количество потокв для пингования
	AddressBufferSize *uint64 // Размер буфера адресов
)

// Создаём указатели на общие обьекты
var (
	IPChannel    chan string
	RawIPChannel chan string
)

// Создаем WaitGroups для отслеживания завершения горутин
var (
	// WorkWG - Группа горутин отвечающих за запуск/завершение сканирующих потоков
	WorkWG sync.WaitGroup

	// PingWG - Группа горутин отвечающих за запуск/завершение пингующих потоков
	PingWG sync.WaitGroup
)

func main() {
	// Получаем флаги
	InitFlags()
	ParseFlags()
	CheckFlags()

	// Создаём сеть
	net := NewNetwork(*InputNet)

	// Подключаем базу данных

	// Инициализируем канал для передачи проверенных адресов
	IPChannel = make(chan string, *AddressBufferSize)

	// Инициализирую канал для передачи непроверенных адресов
	RawIPChannel = make(chan string, *AddressBufferSize)

	// Инициализируем рабочие потоки
	log.Println("Запускаю сканирующие потоки")
	for i := uint64(0); i < *NumberScanThreads; i++ {
		WorkWG.Add(1)
		go ScannerThread(IPChannel)
	}
	log.Println("Запуск сканирующих потоков завершён")

	// Инициализируем пингующие потоки
	log.Println("Запускаю пингующие потоки")
	for i := uint64(0); i < *NumberPingThreads; i++ {
		PingWG.Add(1)
		go PingingThread(RawIPChannel, IPChannel)
	}
	log.Println("Запуск пингующих потоков завершён")

	// Запускаем основную петлю для генерации и передачи адресов
	log.Println("Запускаю генерирующую петлю")
	for !net.ended {
		if !net.IsPrivate() {
			RawIPChannel <- net.String()
		}

		net.Inc()
	}
	log.Println("Генерирующая петля завершила работу")
	log.Println("Ожидаем завершение работы пингующих потоков")

	// Закрываем канал пингующих потоков
	for {
		if len(RawIPChannel) == 0 {
			close(RawIPChannel)
			break
		}
	}

	// Ожидаем завершение работы пингующих потоков
	PingWG.Wait()

	log.Println("Пингующие потоки завершили свою работу")
	log.Println("Ожидаем завершение работы сканирующих потоков")

	// Закрываем канал сканирующих потоков
	for {
		if len(IPChannel) == 0 {
			close(IPChannel)
			break
		}
	}

	// Ожидаем завершение работы сканирующих потоков
	WorkWG.Wait()

	log.Println("Сканирующие потоки завершили свою работу")
}
