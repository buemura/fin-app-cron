package internal

import (
	"fmt"
	"os"
)

var URL string

func LoadEnv() {
	URL = os.Getenv("BACKEND_URL_RESET_EXPENSES")
	fmt.Println(URL)
}