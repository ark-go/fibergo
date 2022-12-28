package nasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// type birds struct {
// 	Url string
// }

type birds []string

func Birds() ([]string, error) {
	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	urll := "https://shibe.online/api/birds"
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

	svc := birds{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if len(svc) > 0 {
		log.Println("Dog: ", svc[0])
		return []string{svc[0]}, nil
	} else {
		return nil, err
	}
}
