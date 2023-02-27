package pkg

import (
	"time"

	"github.com/buemura/golang-cron-jobs/internal"
	"github.com/go-co-op/gocron"
)

func RunCronJob() {
	c := gocron.NewScheduler(time.UTC)

	c.Every(1).Month(28).Do(func() {
		internal.UpdateAllExpenses()
	})

	c.StartBlocking()
}