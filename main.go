package main

import (
	"fmt"
	"os"
	"tsForcast/stock"
)


func main() {
	firm := os.Args[1]
	fmt.Println("Get Stock price for [ ", firm, " ] from Google")
	fmt.Println(stock.GetStockDetails(firm))
}
