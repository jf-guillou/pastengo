#Pastengo
[![Build Status](https://travis-ci.org/jf-guillou/pastengo.svg?branch=master)](https://travis-ci.org/jf-guillou/pastengo) [![GoDoc](https://godoc.org/github.com/jf-guillou/pastengo?status.svg)](https://godoc.org/github.com/jf-guillou/pastengo) [![Go Report Card](https://goreportcard.com/badge/github.com/jf-guillou/pastengo)](https://goreportcard.com/report/github.com/jf-guillou/pastengo) [![MIT
licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/jf-guillou/pastengo/master/LICENSE.md)

Modern self-hosted pastebin service written in Go.

## Motivation
Many Pastebin services exist but all are more complicated than they need to be.
That is why I decided to write a pastebin service in golang.

This was originally a fork from [https://github.com/ewhal/Pastebin](https://github.com/ewhal/Pastebin), 
although the code has severely been modified and functionalities altered.

Supports mysql and sqlite3 database drivers.

![Paste here](https://img.thetabx.net/ZpeZn.jpg)
![Paste content](https://img.thetabx.net/xi3sy.jpg)

## Getting started
### Prerequisities
* [go](https://golang.org/doc/install)
* gcc
* mysql (optionnal)

### Installing

```
go get https://github.com/jf-guillou/pastengo
cd /path/to/install
go build github.com/jf-guillou/pastengo
go generate github.com/jf-guillou/pastengo
```

### Updating
```
go get -u https://github.com/jf-guillou/pastengo
cd /path/to/install
go build github.com/jf-guillou/pastengo
```

### Configuration
Go generate should create a copy of the configuration template.
You should edit ```config.json``` before running pastengo.
```
Address : Base url used for redirections
Port : HTTP listen port
Length : Length of paste ID
DBType : "sqlite3" or "mysql" based on your preferences
DBName : Database sqlite filename or mysql name 
DBUsername : Mysql database username (unused for sqlite)
DBPassword : Mysql database password (unused for sqlite)
```

### Run
Pastengo will automatically create the required SQL tables required.
Running the app is very simple, although we recommand using a process control system like [Supervisord](http://supervisord.org/).
```
./pastengo
```

## FAQ

#### go generate fails
go generate relies on commands that may not be available on your system.
Doing this manually is pretty simple :
* copy folder ```$GOPATH/src/github.com/jf-guillou/pastengo/assets``` to ```/path/to/install/```
* copy file ```$GOPATH/src/github.com/jf-guillou/pastengo/config.example.json``` to ```/path/to/install/config.json```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

