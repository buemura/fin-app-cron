package internal

import (
	"fmt"
	"strings"

	"github.com/buemura/golang-cron-jobs/shared"
)

func UpdateAllExpenses() {
	fmt.Println("PUT request to: ", URL)
	res := shared.PutRequest(URL, strings.NewReader("no data"))
	fmt.Println("Response: ", res)
}
