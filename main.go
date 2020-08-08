package main

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/utility"
)

func main() {

	fmt.Println(utility.GetName())

	myStock := models.Trade{
		Symbol: "MSFT",
		Volume: 10,
		Price:  100.10,
		Buy:    true,
	}

	fmt.Printf("Stock name : %v", myStock.Price)

}
