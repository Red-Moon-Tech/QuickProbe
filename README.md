[![CI сборка](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/go.yml/badge.svg)](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/go.yml)
[![CI сборка](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/makefile.yml/badge.svg)](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/makefile.yml)
# QuickProbe
QuickProbe — это высокопроизводительный сетевой сканер, разработанный на языке Golang, который позволяет эффективно сканировать участки сетей и определять открытые IP-адреса и порты. Программа акцентирует внимание на скорости и параллелизме, что позволяет обрабатывать большие объемы данных за минимальное время.
### Install
____
Build QuickProbe:
```
make build
```
Specify architecture and OS:
```
Example:
make GOOS=windows GOARCH=amd64
```
Clean build files:
```
make clean
```
### Usage
____
Commands description:
```
./quickprobe --help
```
Example of usage:
```
./quickprobe --Network 192.168.1.0/24 --SkipPrivateRange=false 
```
Example of usage #2:
```
./quickprobe --Network ip.txt --Timeout=50 -sT=10 -sP=10 -p=22,80,443,3030
```
