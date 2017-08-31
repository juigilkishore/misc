package stock

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Stock struct {
	Name	string
	Price	int
}

func GetNewStockInstance() *Stock {
	stock := new(Stock)
	return stock
}

func frameURIfromTickerSymbol(firm string) string {
	baseuri := "http://finance.google.com/"
	extension := "finance/info?client=ig&q=NASDAQ%3A"
	return baseuri + extension + firm
}

func getDataFrom(firm string) string {
	firmUri := frameURIfromTickerSymbol(firm)
	fmt.Println(firmUri)
	resp, err := http.Get(firmUri)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	ioutil.Readall(resp.Body)
	return "empty string"
	/*
	body, err := ioutil.Readall(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyStringify := string(body)
	bodyStringifiesLength := len(bodyStringify)
	requiredBodyString := bodyStringify[6:bodyStringifiesLength-3]
	return requiredBodyString
	*/

}

func GetStockDetails(firm string) string {
	return getDataFrom(firm)

}
