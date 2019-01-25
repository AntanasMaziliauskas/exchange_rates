package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AntanasMaziliauskas/exchange_rates/rates"
)

//This function lets me input a currency using flag and prints out the information of that specific currency..
func main() {

	base := flag.String("currency", "USD", "a string")
	startdate := flag.String("startdate", "2018-02-02", "a string")
	enddate := flag.String("enddate", "2018-02-06", "a string")
	flag.Parse()

	URLNameStart := fmt.Sprintf(rates.URL, *startdate, *base)
	URLNameEnd := fmt.Sprintf(rates.URL, *enddate, *base)

	if errStart, exratesStart := rates.FromURL(URLNameStart); errStart != nil {
		fmt.Println(errStart.Error())
		log.Fatal("Error")
	}
	errEnd, exratesEnd := rates.FromURL(URLNameEnd)
	if errEnd != nil {
		fmt.Println(errStart.Error())
		log.Fatal("Error")
	}

	fmt.Printf("Base value: %s \n", exratesStart.Base)
	fmt.Printf("Start Date: %s \n", exratesStart.Date)
	fmt.Printf("End Date: %s \n", exratesEnd.Date)
	for k, v := range exratesStart.Rates {
		fmt.Printf("Currency: %s Difference: %5.2f%%\n", k, rates.Percentage(v, exratesEnd.Rates[k]))
	}
}
