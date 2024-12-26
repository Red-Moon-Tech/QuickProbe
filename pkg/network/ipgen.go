package network

import (
	"QuickProbe/pkg/argflags"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Обьявляем приватные диапазоны
var (
	privateRangeClassA = NewNetwork("10.0.0.0/8")      // Класс A
	privateRangeClassB = NewNetwork("172.16.0.0/12")   // Класс B
	privateRangeClassC = NewNetwork("192.168.0.0/16")  // Класс C
	privateRangeShared = NewNetwork("100.64.0.0/10")   // Shared Address Space (RFC 6598)
	linkLocalRange     = NewNetwork("127.0.0.0/8")     // Локальный адрес (loopback)
	specialRange       = NewNetwork("0.0.0.0-1.0.0.0") // Специальные адреса
)

type Network struct {
	hostMin        uint
	hostMax        uint
	totalAddresses uint
	currentAddress uint
	Ended          bool
}

// IsPrivate проверяет относится ли текущий адрес (Network.currentAddress) к приватным диапазонам
func (net *Network) IsPrivate() (bool, *Network) {
	switch {
	case net.IsPartOfNetwork(privateRangeClassA):
		return true, privateRangeClassA
	case net.IsPartOfNetwork(privateRangeClassB):
		return true, privateRangeClassB
	case net.IsPartOfNetwork(privateRangeClassC):
		return true, privateRangeClassC
	case net.IsPartOfNetwork(privateRangeShared):
		return true, privateRangeShared
	case net.IsPartOfNetwork(linkLocalRange):
		return true, linkLocalRange
	case net.IsPartOfNetwork(specialRange):
		return true, specialRange

	default:
		return false, nil

	}
}

// IsPartOfNetwork проверяет принадлежит ли текущий адрес (Network.currentAddress) к указаной сети
func (net *Network) IsPartOfNetwork(AnotherNetwork *Network) bool {
	if net.currentAddress < AnotherNetwork.hostMax && net.currentAddress >= AnotherNetwork.hostMin {
		return true
	} else {
		return false
	}
}

// Inc функция инкрементирует текущий адрес
func (net *Network) Inc() {
	net.currentAddress += 1

	// Пропускаем приватный диапазон
	if *argflags.SkipPrivateRange {
		private, privateNet := net.IsPrivate()

		if private {
			net.currentAddress = privateNet.hostMax + 1
		}
	}

	if *argflags.SkipAddressRange != "None" {
		SkipAddr := NewNetwork(*argflags.SkipAddressRange)
		if net.IsPartOfNetwork(SkipAddr) {
			net.currentAddress = SkipAddr.hostMax + 1
		}
	}

	// Проверка на достижение края диапазона
	if net.currentAddress > net.hostMax {
		net.Ended = true
	}
}

// Возвращает текущий IP ввиде строки
func (net *Network) String() string {
	// Получаем октеты
	firstOctet := net.currentAddress >> 24 & 0xFF
	secondOctet := net.currentAddress >> 16 & 0xFF
	thirdOctet := net.currentAddress >> 8 & 0xFF
	fourthOctet := net.currentAddress & 0xFF

	return fmt.Sprintf("%d.%d.%d.%d", firstOctet, secondOctet, thirdOctet, fourthOctet)
}

// NewNetwork - конструктор сети
func NewNetwork(address string) *Network {
	// Определяем переменные для полей
	var (
		hostMin        uint
		hostMax        uint
		bitmask        uint
		totalAddresses uint
	)

	// Определяем тип адреса
	if strings.Contains(address, "/") {
		// Отделяем IP от маски
		splitAddress := strings.Split(address, "/")
		ip := splitAddress[0]
		bm, err := strconv.Atoi(splitAddress[1])
		if err != nil {
			panic(err)
		}
		bitmask = uint(bm)

		// Получаем количество адресов в сети
		totalAddresses = uint(math.Pow(2, float64(32-bitmask)))

		// Преобразуем IP в число
		intIP := ipToInt(ip)

		// Получаем минимальный и максимальный адрес
		hostMin = intIP & (intIP ^ uint(math.Pow(2, float64(32-bitmask))-1))
		hostMax = hostMin + totalAddresses
	} else if strings.Contains(address, "-") {
		// Отделяем начальный и конечный адрес
		addresses := strings.Split(address, "-")

		// Преобразуем IP в числа
		initialIP := ipToInt(addresses[0])
		finalIP := ipToInt(addresses[1])

		// Устанавлиаем минимальный и максимальный адрес
		hostMin = initialIP
		hostMax = finalIP

		// Устанавлиаем количество адресов в диапазоне
		totalAddresses = finalIP - initialIP
	}

	return &Network{
		hostMin:        hostMin,
		hostMax:        hostMax,
		totalAddresses: totalAddresses,
		currentAddress: hostMin,
		Ended:          false,
	}
}
