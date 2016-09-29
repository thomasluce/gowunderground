package gowunderground

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// This is a very simple WeatherUnderground (wunderground.com) client. I only
// care about one aspect of their api, so I only implement that one. If you
// want to add other parts of the API, pull requests are welcome.

// WundergroundChance is a percentage chance of something.
type WundergroundChance struct {
	Percentage string `json:"percentage"`
}

// WundergroundChanceOf is a collection of percentage chances of cloudy and
// partly cloudy days
type WundergroundChanceOf struct {
	Chanceofcloudyday       WundergroundChance `json:"chanceofcloudyday"`
	Chanceofpartlycloudyday WundergroundChance `json:"chanceofpartlycloudyday"`
}

// WundergroundTrip is a trip record, which contains, among other things, the cloud chances for the
// time-period.
type WundergroundTrip struct {
	ChanceOf WundergroundChanceOf `json:"chance_of"`
}

// WundergroundResponse is a response from WeatherUnderground
type WundergroundResponse struct {
	Trip WundergroundTrip `json:"trip"`
}

// Planner returns a `plan` response from Weather Underground (or an error if
// something went wrong) based on the api key, start and end times, and
// location.
func Planner(apiKey string, startDate time.Time, endDate time.Time, lat, lon float64) (WundergroundResponse, error) {
	url := fmt.Sprintf("http://api.wunderground.com/api/%s/planner_%02d%02d%02d%02d/q/%f,%f.json", apiKey, startDate.Month(), startDate.Day(), endDate.Month(), endDate.Day(), lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return WundergroundResponse{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var results WundergroundResponse
	err = json.Unmarshal(body, &results)
	if err != nil {
		return WundergroundResponse{}, err
	}

	return results, nil

}
