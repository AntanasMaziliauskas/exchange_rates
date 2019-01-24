package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ExRates struct { // comment
	Rates map[string]float32 `json:"rates"` //map
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

func main() {
	//reader := bufio.newReader(os.Stdin)
	//var opt string
	var content []byte
	fmt.Println("Press 'f' if you want to use File or press something else if you want to use URL...")
	var opt string
	fmt.Scanln(&opt)
	if opt == "f" {
		DocName := "rates.json"
		content := FromDocument(DocName)
	} else {
		URLName := "https://api.exchangeratesapi.io/latest?base=USD"
		content := FromURL(URLName)
	}
	//fmt.Println(first)
	//opt, _ := reader.readString("\n")
	/*switch opt {
	case "F":
		DocName := "rates.json"
		content := FromDocument(DocName)
	case "U":
		URLName := "https://api.exchangeratesapi.io/latest?base=USD"
		content := FromURL(URLName)
	default:
		fmt.Println("Error: Wrong button pressed.")
		break
	}*/
	//DocName := "rates.json"
	//content := FromDocument(DocName)
	exrates := ExRates{}
	// json.Unmarshal(content, &friends)
	err2 := json.Unmarshal(content, &exrates)
	if err2 != nil {
		fmt.Println("Error JSON Unmarshalling")
		fmt.Println(err2.Error())
	}
	fmt.Printf("Base value: %s \n", exrates.Base)
	fmt.Printf("Date: %s \n", exrates.Date)
	//	if len(exrates.Rates) != 0 {aaasdas
	//	number := exrates.Rates["HUF"]
	//	fmt.Printf("Test: %f \n", number)
	//fmt.Printf("Rates: %v \n", exrates.Rates)
	//}
	for k, v := range exrates.Rates { // map
		fmt.Printf("Currency: %s Value: \t%.2f\n", k, v)
	}

}

func FromDocument(DocName string) []byte {
	content, err := ioutil.ReadFile(DocName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return content
}

func FromURL(URLName string) []byte {
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
	return body
}
