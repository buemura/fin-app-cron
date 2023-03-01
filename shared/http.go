package shared

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func PutRequest(url string, data io.Reader) (string, error) {
	client := &http.Client{}
	var resp *http.Response
	var bodyBytes []byte

	req, err := http.NewRequest(http.MethodPut, url, data)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}