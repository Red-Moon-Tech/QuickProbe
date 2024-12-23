[![CI сборка](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/go.yml?query=workflow%3AGo)
[![CI сборка](https://github.com/Red-Moon-Tech/QuickProbe/actions/workflows/makefile.yml/badge.svg?branch=master)](https://github.com/Red-Moon-Tech/QuickProbe/actions?query=workflow%3AMakefile%20CI)
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
