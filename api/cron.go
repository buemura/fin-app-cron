package handler

import (
	"net/http"

	"github.com/buemura/golang-cron-jobs/internal"
)

func init() {
	internal.LoadEnv()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	internal.UpdateAllExpenses()
}
