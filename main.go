package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ExRates struct { // comment
	Rates map[string]float32 `json:"rates"` //map
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

func main() {

	content, err := ioutil.ReadFile("rates.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	exrates := ExRates{}
	// json.Unmarshal(content, &friends)
	err2 := json.Unmarshal(content, &exrates)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	fmt.Printf("Base value: %s \n", exrates.Base)
	fmt.Printf("Date: %s \n", exrates.Date)
	//	if len(exrates.Rates) != 0 {
	//	number := exrates.Rates["HUF"]
	//	fmt.Printf("Test: %f \n", number)
	//fmt.Printf("Rates: %v \n", exrates.Rates)
	//}
	for k, v := range exrates.Rates { // map
		fmt.Printf("Currency: %s Value: \t%.2f\n", k, v)
	}

}
