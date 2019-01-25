package rates

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

//ExRates structure used to unmarshal JSON
type ExRates struct {
	Rates map[string]float32 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

const URL = "https://api.exchangeratesapi.io/%s?base=%s"

// Function FromURL takes URL address of a JSON, unmarshals JSON and return the data
func FromURL(URLName string) (ExRates, error) {
	spaceClient := http.Client{
		Timeout: time.Second * 10, // 2 secs
	}
	exrates := ExRates{}
	req, err := http.NewRequest(http.MethodGet, URLName, nil)
	if err != nil {
		return exrates, err
		//fmt.Println(err.Error())
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return exrates, getErr
		//log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return exrates, readErr
		//	log.Fatal(readErr)
	}
	//exrates := ExRates{}
	// json.Unmarshal(content, &friends)
	if err := json.Unmarshal(body, &exrates); err != nil {
		return exrates, err
		//	fmt.Println("Error JSON Unmarshalling")
		//	fmt.Println(err.Error())
		//	log.Fatal("error")
	}
	return exrates, nil
}

// Function Percentage count the percentage of difference between 'start' and 'end' variables
func Percentage(start, end float64) float64 {
	percentDiff := (end - start) * 100 / end
	percentDiff = math.Round(percentDiff*10000) / 10000
	return percentDiff
}
