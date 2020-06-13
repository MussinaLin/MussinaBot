package Bitfinex

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
)

type MussinaOrder struct{
	Amount float64
	Rate float64
	Period int64
}



func GenOrders(availableBalance float64, maxSignleOrderAmount float64, minLoan float64,
				left float64) *[]MussinaOrder{
	log.Println(fmt.Sprintf("[GenOrders] availableBalance:%f, maxSignleOrderAmount:%f, " +
		"minLoan:%f left:%f", availableBalance, maxSignleOrderAmount, minLoan, left))

	availableBalance -= left
	orders := make([]MussinaOrder, 0)
	num := 0
	for availableBalance > 0 && availableBalance >= minLoan{
		num++
		if availableBalance < maxSignleOrderAmount{ // last order
			orders= append(orders, MussinaOrder{Amount:availableBalance, Rate:0, Period:2})
			availableBalance = 0
		}else{
			orders= append(orders, MussinaOrder{Amount:maxSignleOrderAmount, Rate:0, Period:2})
			availableBalance -= maxSignleOrderAmount
		}
	}
	log.Println(fmt.Sprintf("[GenOrders] gen %d orders", len(orders)))
	return &orders
}

func AssignRate(FRR float64, increaseRate float64, orders *[]MussinaOrder) *[]MussinaOrder{
	log.Println(fmt.Sprintf("[AssignRate] FRR:%f, increaseRate:%f", FRR, increaseRate))
	for index, _ := range *orders{
		(*orders)[index].Rate = FRR + FRR * increaseRate * 0.01 * float64(index)
	}
	return orders
}

func SubmitOrders(orders *[]MussinaOrder){
	for _, order := range *orders{
		noti,err := bfRestClient.Funding.SubmitOffer(&bitfinex.FundingOfferRequest{
			Type:"LIMIT",
			Symbol:"fUSD",
			Amount:order.Amount,
			Rate:order.Rate,
			Period:order.Period,
			Hidden:false,
		})
		if err != nil{
			log.Println("[ERROR] ", err.Error())
		}
		log.Println(noti)
	}
	
}