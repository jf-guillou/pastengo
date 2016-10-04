#Pastengo
[![Build Status](https://travis-ci.org/jf-guillou/pastengo.svg?branch=master)](https://travis-ci.org/jf-guillou/pastengo) [![GoDoc](https://godoc.org/github.com/jf-guillou/pastengo?status.svg)](https://godoc.org/github.com/jf-guillou/pastengo) [![Go Report Card](https://goreportcard.com/badge/github.com/jf-guillou/pastengo)](https://goreportcard.com/report/github.com/jf-guillou/pastengo) [![MIT
licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jf-guillou/pastengo/master/LICENSE.md)

Modern self-hosted pastebin service written in Go.

## Motivation
Many Pastebin services exist but all are more complicated than they need to be.
That is why I decided to write a pastebin service in golang.

This was originally a fork from [https://github.com/ewhal/Pastebin](https://github.com/ewhal/Pastebin), 
although the code has severely been modified and functionalities altered

Supports mysql and sqlite3 database drivers

![Paste here](https://img.thetabx.net/ZpeZn.jpg)
![Paste content](https://img.thetabx.net/xi3sy.jpg)

## Getting started
### Prerequisities
* go
* mysql (optionnal)

### Installing

* go get https://github.com/jf-guillou/pastengo
* cp config.example.json config.json
* nano config.json
* Configure address, port and database details

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

