package rates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
func FromURL(URLName string) ExRates {
	spaceClient := http.Client{
		Timeout: time.Second * 10, // 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, URLName, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	exrates := ExRates{}
	// json.Unmarshal(content, &friends)
	if err := json.Unmarshal(body, &exrates); err != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err.Error())
		log.Fatal("error")
	}
	return exrates
}

// Function Percentage count the percentage of difference between 'start' and 'end' variables
func Percentage(start, end float32) float32 {
	percentDiff := (end - start) * 100 / end
	return percentDiff
}
