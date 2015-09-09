package main

import (
	"log"

	"github.com/kohkimakimoto/loglv"
)

func main() {
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
