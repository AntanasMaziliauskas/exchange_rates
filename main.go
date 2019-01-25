package main

import (
	"encoding/json"
	"flag"
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

//This function lets me input a currency using flag and prints out the information of that specific currency
func main() {

	base := flag.String("currency", "USD", "a string")
	startdate := flag.String("startdate", "2018-02-02", "a string")
	enddate := flag.String("enddate", "2018-02-06", "a string")
	flag.Parse()

	URLNameStart := fmt.Sprintf(URL, *startdate, *base)
	URLNameEnd := fmt.Sprintf(URL, *enddate, *base)

	exratesStart := FromURL(URLNameStart)
	exratesEnd := FromURL(URLNameEnd)

	fmt.Printf("Base value: %s \n", exratesStart.Base)
	fmt.Printf("Start Date: %s \n", exratesStart.Date)
	fmt.Printf("End Date: %s \n", exratesEnd.Date)
	for k, v := range exratesStart.Rates {
		fmt.Printf("Currency: %s Difference: %5.2f%%\n", k, Percentage(v, exratesEnd.Rates[k]))
	}
}

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

func Percentage(start, end float32) float32 {
	percentDiff := (end - start) * 100 / end
	return percentDiff
}
