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
	Rates map[string]float32 `json:"rates"` //map
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

func main() {
	const URL = "https://api.exchangeratesapi.io/latest?base="
	wordPtr := flag.String("currency", "USD", "a string")
	flag.Parse()
	//	fmt.Println("word:", *wordPtr)
	//Sujungiam URL su base
	// var buffer bytes.Buffer
	// buffer.WriteString(URL)
	// buffer.WriteString(*wordPtr)
	// URLName := buffer.String()
	URLName := fmt.Sprintf("%s%s", URL, *wordPtr)
	// fmt.Println(buffer.String())
	//var content []byte
	exrates := FromURL(URLName)
	//reader := bufio.newReader(os.Stdin)
	//var opt string

	/*fmt.Println("Press 'f' if you want to use File or press something else if you want to use URL...")
	var opt string
	fmt.Scanln(&opt)
	if opt == "f" {
		//	DocName := "rates.json"
		//content = FromDocument(DocName)
	} else {
		URLName := "https://api.exchangeratesapi.io/latest?base=USD"
		content = FromURL(URLName)
	}*/
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

	fmt.Printf("Base value: %s \n", exrates.Base)
	fmt.Printf("Date: %s \n", exrates.Date)
	//	if len(exrates.Rates) != 0 {aaasdas
	//	number := exrates.Rates["HUF"]
	//	fmt.Printf("Test: %f \n", number)kk
	//fmt.Printf("Rates: %v \n", exrates.Rates)
	//}
	for k, v := range exrates.Rates { // map
		fmt.Printf("Currency: %s Value: %8.2f\n", k, v)
	}

	//URL := append("kaboom", *wordPrt)
	//fmt.Println("Test: %v", URL)
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
	// return exrates.Base, exrates.Date, exrates.Rates
	return exrates
}
