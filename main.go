package main

import "fmt"
import "tsForcast/stock"

func main() {
    fmt.Println("Get Stock price from Google")
    s := stock.GetNewStockInstance()
    fmt.Println(s)
}
