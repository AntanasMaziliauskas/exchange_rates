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

//ExRates
type ExRates struct {
	Rates map[string]map[string]float32 `json:"rates"` //map
	//Base      string `json:"base"`
	EndDate   string `json:"end_at"`
	StartDate string `json:"start_at"`
}

const URL = "https://api.exchangeratesapi.io/history?start_at=%s&end_at=%s"

func main() {

	startdate := flag.String("start", "2018-01-02", "date")
	enddate := flag.String("end", "2018-01-03", "date")
	flag.Parse()

	URLName := fmt.Sprintf(URL, *startdate, *enddate)

	exrates := FromURL(URLName)

	fmt.Printf("Start date: %s \n", exrates.StartDate)
	fmt.Printf("End date: %s \n", exrates.EndDate)
	fmt.Printf("Test: %s \n", exrates.Days.Rates)

	for k, v := range exrates.Days.Rates { // map
		fmt.Printf("Currency: %s Value: %s\n", k, v)
	}

}

// func FromURL(URLName string) (string, string, map[string]float32) {
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
	err2 := json.Unmarshal(body, &exrates)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}

	return exrates
}
