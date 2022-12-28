package nasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type fox struct {
	Link string
}

func Fox() ([]string, error) {
	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	urll := "https://randomfox.ca/floof/"
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

	svc := fox{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	log.Println("Fox: ", svc.Link)
	return []string{svc.Link}, nil
}
