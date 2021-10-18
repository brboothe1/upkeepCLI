package main

import (
	"github.com/robfig/cron"
)

// funcScheduler implements a cron that schedules reminders throughout the day

func funcScheduler() {
	c := cron.New()

	c.AddFunc("0 8 * * *", createCommitmentFile)
	c.AddFunc("30 8,14,17 * * *", createTextMessage)
	c.AddFunc("0 22 * * *", recordDailyReflection)

	c.Run()
}
