package nasa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&count=1
type zap struct {
	Copyright string
	Date      string
	// описание
	Explanation    string
	Hdurl          string
	MediaType      string
	ServiceVersion string
	Title          string
	Url            string
}

func LoadNasa() ([]zap, error) {

	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, "https://api.nasa.gov/planetary/apod?api_key="+os.Getenv("Nasa_Key")+"&count=1", nil)
	if err != nil {
		return nil, err
	}

	res, getErr := testClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	svc := []zap{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	//log.Printf("%+v", svc)
	return svc, nil
}

/*
[{"copyright":"U. Chicago",
"date":"1998-02-25",
"explanation":"Вы здесь. Оранжевая точка на приведенном выше рисунке в искусственном цвете представляет текущее положение Солнца.",
"hdurl":"https://apod.nasa.gov/apod/image/9802/sunjourney_pf_big.gif",
"media_type":"image",
"service_version":"v1",
"title":"Солнечное соседство",
"url ":"https://apod.nasa.gov/apod/image/9802/sunjourney_pf.jpg"}]
*/
