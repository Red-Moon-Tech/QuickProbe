package main

import (
	"flag"
	"log"
)

func InitFlags() {
	InputNet = flag.String("Network", "None", "Сеть для сканирования")
	NumberScanThreads = flag.Uint64("NumberScanThreads", 5, "Количество сканирующих потоков")
	NumberPingThreads = flag.Uint64("NumberPingThreads", 5, "Количество пингующих потоков")
	AddressBufferSize = flag.Uint64("AddressBufferSize", 0, "Размеров буфера адресов")
}

func ParseFlags() {
	flag.Parse()

	if *AddressBufferSize == 0 {
		*AddressBufferSize = *NumberScanThreads
	}
}

func CheckFlags() {
	if *InputNet == "None" {
		log.Fatal("Не указана сеть для сканирования, используйте флаг: --Network")
	}
	if *AddressBufferSize < *NumberScanThreads {
		log.Println("WARNING: Рекомендуется устанавливать размер буфера адресов неменее количества сканирующих потоков")
	}
}
