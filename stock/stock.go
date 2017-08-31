package stock

import "net/http"

type Stock struct {
	Name	string
	Price	int
}

func GetNewStockInstance() *Stock {
	stock := new(Stock)
	return stock
}

func frameURIfromTickerSymbol(tickerSymbol string) string {
	// do something
}

func GetDataFrom(tickerSymbol string) {
	uri := frameURIfromTickerSymbol(tickerSymbol)
	resp, _ := http.Get(uri)
}
