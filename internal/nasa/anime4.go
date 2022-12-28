package nasa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type anime4 struct {
	Images []struct {
		Url string
	}
}

func Anime4() ([]string, error) {
	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	//! пересмотреть генерацию в init !!! чтоб использовать это:
	//!    generator := rand.New(rand.NewSource(time.Now().UnixNano())) ;
	//!    generator.Intn...
	//generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	kategory := []string{
		"ass", // NSFW
		"hentai",
		"milf",
		"oral",
		"paizuri",
		"ecchi",
		"ero",
	}
	/* = versatile
	maid
	maifu
	marin-kitagawa
	mori-calliope
	raiden-shogun
	oppai
	selfies
	uniform
	*/

	urll := "https://api.waifu.im/search/?included_tags=" + kategory[rand.Intn(len(kategory))]

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
	svc := anime4{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if len(svc.Images) > 0 {
		log.Println("Anime4: ", svc.Images[0].Url)
		return []string{svc.Images[0].Url}, nil
	} else {
		return nil, errors.New("Anime4 пусто")
	}
}
