package cron

import (
	"github.com/robfig/cron"
)

func EveryMinute(a func()) {
	c := cron.New()
	// do query every 30 seconds
	c.AddFunc("*/0 * * * *", a)
	c.Start()
	select {}
}
