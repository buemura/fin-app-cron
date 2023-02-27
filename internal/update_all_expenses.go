package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func putRequest(url string, data io.Reader) string {
	client := &http.Client{}
	var resp *http.Response
	var bodyBytes []byte

	req, err := http.NewRequest(http.MethodPut, url, data)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    bodyString := string(bodyBytes)

	return bodyString
}

func UpdateAllExpenses() {
	fmt.Println("PUT request to: ", URL)
	res := putRequest(URL, strings.NewReader("no data"))
	fmt.Println("Response: ", res)
}
