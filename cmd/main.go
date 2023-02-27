package main

import (
	"github.com/buemura/golang-cron-jobs/pkg"
)

func init() {
	pkg.LoadEnv()
}

func main() {
	pkg.RunCronJob()
}