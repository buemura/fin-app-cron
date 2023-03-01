package internal

import (
	"fmt"

	"github.com/buemura/golang-cron-jobs/shared"
)

func UpdateAllExpenses() (error) {
	fmt.Println("PUT request to: ", URL)
	
	res, err := shared.PutRequest(URL, nil)
	if err != nil {
		return err
	}

	fmt.Println("Response: ", res)
	return nil
}
