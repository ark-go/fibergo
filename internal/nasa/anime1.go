package nasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type anime1 struct {
	Data struct {
		Images struct {
			Jpg struct {
				LargeImageUrl string `json:"large_image_url"`
			}
		}
	}
	Title string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Anime1() ([]string, error) {
	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	kategory := []string{
		"manga", //  404
		"anime", //  45

	}
	urll := "https://api.jikan.moe/v4/random/" + kategory[rand.Intn(len(kategory))]

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
	//log.Printf("%+v", string(body))
	svc := anime1{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	log.Println("Anime1: ", svc.Data.Images.Jpg.LargeImageUrl)
	return []string{svc.Data.Images.Jpg.LargeImageUrl}, nil
}
