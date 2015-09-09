# loglv

[![Build Status](https://travis-ci.org/kohkimakimoto/loglv.svg)](https://travis-ci.org/kohkimakimoto/loglv)

A package supports leveled logging for Go standard `log` package.

## Installation

Use `go get` command.

```
$ go get github.com/kohkimakimoto/loglv
```

## Usage

```go
package main

import (
	"log"

	"github.com/kohkimakimoto/loglv"
)

func main() {
	// At first, You should run Init method that replaces output of Go standard log package.
	loglv.Init()

	log.Println("standard log message. this is outputted!")
	// ouptput: 2015/09/09 11:18:36 standard log message. this is outputted!

	if loglv.IsDebug() {
		log.Println("debug log message. this is not outputted, because default log level is info!")
	}

	loglv.SetLv(loglv.LvDebug)
	if loglv.IsDebug() {
		log.Println("debug log message. this is outputted!")
	}
	// output: 2015/09/09 11:18:36 debug log message. this is outputted!

	loglv.SetLv(loglv.LvQuiet)
	log.Println("standard log message. but this is not outputted, because log level is quiet!")
}
```

## Author

Kohki Makimoto <kohki.makimoto@gmail.com>

## License

MIT license.
