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

// https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&count=1
type epic struct {
	Identifier string
	Image      string
	Url        string
	// описание
	Caption string
	// CentroidCoordinates string // {"lat":-25.664062,"lon":163.696289}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func ranDate() time.Time {

	min := time.Date(2022, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2022, 12, 15, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
func LoadEpic() ([]epic, error) {

	testClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	//dat := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Local)
	dat := ranDate()
	dateStr := dat.Format("<b>02.01.2006</b>")
	//level := "natural" // "enhanced"
	level := "enhanced"
	urll := fmt.Sprintf("https://epic.gsfc.nasa.gov/api/%s/date/%d-%02d-%02d", level, dat.Year(), dat.Month(), dat.Day())
	log.Printf("Epic urll> %+v", urll)
	req, err := http.NewRequest(http.MethodGet, urll, nil)
	if err != nil {
		return nil, err
	}

	res, getErr := testClient.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("Epic Non-OK HTTP status:", res.StatusCode)

		// You may read / inspect response body
		return nil, getErr
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}

	svc := []epic{}
	jsonErr := json.Unmarshal(body, &svc)
	if jsonErr != nil {
		return nil, jsonErr
	}
	//	log.Printf("Epic > %+v", svc)
	for i, v := range svc {
		urlImage := fmt.Sprintf("https://epic.gsfc.nasa.gov/archive/%s/%d/%02d/%02d/png/%s.png", level, dat.Year(), dat.Month(), dat.Day(), v.Image)
		//	log.Printf("ЖЖЖЖЖЖ> %+v\n", urlImage)
		svc[i].Url = urlImage
		svc[i].Caption = dateStr //+ svc[i].Caption

	}

	//log.Printf("%+v", svc)
	return svc, nil
}

/*
{
"identifier":"20221221001752",
"caption":"This image was taken by NASA's EPIC camera onboard the NOAA DSCOVR spacecraft",
"image":"epic_1b_20221221001752",
"version":"03",
"centroid_coordinates":{"lat":-25.664062,"lon":163.696289},
                        "dscovr_j2000_position":{"x":-297071.683483,"y":-1316902.874094,"z":-651446.5875},
						"lunar_j2000_position":{"x":-208322.388447,"y":-275380.726412,"z":-127687.55554},
						"sun_j2000_position":{"x":-3154516.918034,"y":-135009345.027998,"z":-58525488.599999},
						"attitude_quaternions":{"q0":-0.49518,"q1":-0.57886,"q2":0.25659,"q3":0.59488},
						"date":"2022-12-21 00:13:03",
						"coords":{"centroid_coordinates":{"lat":-25.664062,"lon":163.696289},
						           "dscovr_j2000_position":{"x":-297071.683483,"y":-1316902.874094,"z":-651446.5875},
								   "lunar_j2000_position":{"x":-208322.388447,"y":-275380.726412,"z":-127687.55554},
								   "sun_j2000_position":{"x":-3154516.918034,"y":-135009345.027998,"z":-58525488.599999},
"attitude_quaternions":{"q0":-0.49518,"q1":-0.57886,"q2":0.25659,"q3":0.59488}}},*/
