package internal

import (
	"fmt"

	"github.com/buemura/golang-cron-jobs/shared"
)

func UpdateAllExpenses() {
	fmt.Println("PUT request to: ", URL)
	res := shared.PutRequest(URL, nil)
	fmt.Println("Response: ", res)
}
