package network

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Обьявляем приватные диапазоны
var (
	firstPrivateRange  = NewNetwork("10.0.0.0/8")
	secondPrivateRange = NewNetwork("100.64.0.0/10")
	thirdPrivateRange  = NewNetwork("172.16.0.0/12")
	fourthPrivateRange = NewNetwork("192.168.0.0/16")
	linkLocalRange     = NewNetwork("127.0.0.0/8")
)

type Network struct {
	hostMin        uint
	hostMax        uint
	bitmask        uint
	totalAddresses uint
	currentAddress uint
	Ended          bool
}

// IsPrivate проверяет относится ли текущий адрес (Network.currentAddress) к приватным диапазонам
func (net *Network) IsPrivate() bool {
	if net.IsPartOfNetwork(firstPrivateRange) {
		return true
	} else if net.IsPartOfNetwork(secondPrivateRange) {
		return true
	} else if net.IsPartOfNetwork(thirdPrivateRange) {
		return true
	} else if net.IsPartOfNetwork(fourthPrivateRange) {
		return true
	} else if net.IsPartOfNetwork(linkLocalRange) {
		return true
	}

	return false
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

	// Проверка на достижение края диапазона
	if net.currentAddress == net.hostMax {
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

func NewNetwork(address string) *Network {
	// Отделяем IP от маски
	splitAddress := strings.Split(address, "/")
	ip := splitAddress[0]
	bitmask, err := strconv.Atoi(splitAddress[1])
	if err != nil {
		panic(err)
	}

	// Получаем количество адресов в сети
	totalAddresses := uint(math.Pow(2, float64(32-bitmask)))

	// Преобразуем IP в число
	splitIP := strings.Split(ip, ".")
	var intIP uint
	for i := 0; i < 4; i++ {
		cnt, err := strconv.Atoi(splitIP[i])
		if err != nil {
			panic(err)
		}
		intIP += uint(cnt) * uint(math.Pow(2, float64(24-i*8)))
	}

	// Получаем минимальный и максимальный адрес
	hostMin := intIP & (intIP ^ uint(math.Pow(2, float64(32-bitmask))-1))
	hostMax := hostMin + totalAddresses

	return &Network{
		hostMin:        hostMin,
		hostMax:        hostMax,
		bitmask:        uint(bitmask),
		totalAddresses: totalAddresses,
		currentAddress: hostMin,
		Ended:          false,
	}
}
