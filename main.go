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

//This function lets me input a currency using flag and prints out the information of that specific currency
func main() {

	const URL = "https://api.exchangeratesapi.io/latest?base="

	wordPtr := flag.String("currency", "USD", "a string")
	flag.Parse()

	URLName := fmt.Sprintf("%s%s", URL, *wordPtr)

	exrates := FromURL(URLName)

	fmt.Printf("Base value: %s \n", exrates.Base)
	fmt.Printf("Date: %s \n", exrates.Date)
	for k, v := range exrates.Rates {
		fmt.Printf("Currency: %s Value: %8.2f\n", k, v)
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
	err2 := json.Unmarshal(body, &exrates)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	return exrates
}
