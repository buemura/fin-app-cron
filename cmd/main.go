package main

import (
	"github.com/buemura/golang-cron-jobs/internal"
	"github.com/buemura/golang-cron-jobs/pkg"
)

func init() {
	internal.LoadEnv()
}

func main() {
	pkg.RunCronJob()
}