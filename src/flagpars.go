package main

import (
	"flag"
	"log"
)

func InitFlags() {
	log.Println("Инициализирую флаги")
	InputNet = flag.String("Network", "None", "Сеть для сканирования")
	NumberScanThreads = flag.Uint64("NumberScanThreads", 500, "Количество сканирующих потоков")
	NumberPingThreads = flag.Uint64("NumberPingThreads", 100, "Количество пингующих потоков")
	AddressBufferSize = flag.Uint64("AddressBufferSize", 0, "Размеров буфера адресов")
	log.Println("Инициализация флагов завершена")
}

func ParseFlags() {
	log.Println("Получаю флаги")
	flag.Parse()

	if *AddressBufferSize == 0 {
		*AddressBufferSize = *NumberScanThreads
	}

	log.Println("Получение флагов завершено")
}

func CheckFlags() {
	log.Println("Проверяю флаги")
	if *InputNet == "None" {
		log.Fatal("Не указана сеть для сканирования, используйте флаг: --Network")
	}
	if *AddressBufferSize < *NumberScanThreads {
		log.Println("WARNING: Рекомендуется устанавливать размер буфера адресов неменее количества сканирующих потоков")
	}
	log.Println("Проверка флагов завершена")
}
