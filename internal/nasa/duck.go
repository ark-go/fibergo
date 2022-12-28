package nasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type duck struct {
	Url string
}

func Duck() ([]string, error) {
	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	urll := "https://random-d.uk/api/v2/random"
	req, err := http.NewRequest(http.MethodGet, urll, nil)
	if err != nil {
		return nil, err
	}

	res, getErr := testClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Fox Non-OK HTTP status:", res.StatusCode)

		// You may read / inspect response body
		return nil, getErr
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	svc := duck{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	log.Println("Duck: ", svc.Url)
	return []string{svc.Url}, nil
}
