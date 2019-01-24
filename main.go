package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ExRates struct {
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
	number := exrates.Rates["HUF"]
	fmt.Printf("Test: %s \n", number)
	fmt.Printf("Rates: %s \n", exrates.Rates)
	//	for _, x := range ExRates { // map
	//		fmt.Printf("%s \n", x.Rates)
	//	}

}
