package cliargs

import (
	"flag"
	"log"
)

var DbPath	string
var HostIP	string
var BaseTopic	string
var ServeDir	string

func InitArgs() {
	flag.StringVar(&DbPath, "d", "./data.db", "Database path")
	flag.StringVar(&HostIP, "h", "127.0.0.1", "Mqtt host IP (broker IP)")
	flag.StringVar(&BaseTopic, "t", "/#", "Base topic, don't include '#' or '*'")
	flag.StringVar(&ServeDir, "s", "frontend/dist/", "Serve directory, the directory served as the frontend.")
	flag.Parse()
	log.Println("DB Path: ", DbPath)
	log.Println("Host IP: ", HostIP)
	log.Println("Base topic: ", BaseTopic)
	log.Println("Serve dir: ", ServeDir)
}
