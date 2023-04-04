package main

import (
	"fmt"
	"log"

	"github.com/laenix/webhook/config"
	"github.com/laenix/webhook/pkg/cron"
)

func main() {
	log.Println("[+] start devops webhook service.")
	config.InitConfig()
	cron.EveryMinute(func() { fmt.Println("hello world") })
}
