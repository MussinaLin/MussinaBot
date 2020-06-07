package Bitfinex

type MussinaOrder struct{
	Amount float64
	Rate float64
	Period int32
}



func GenOrders(availableBalance float64, maxSignleOrderAmount float64, minLoan float64, left float64) int{
	availableBalance -= left
	orders := make([]MussinaOrder, 0)
	num := 0
	for availableBalance > 0 && availableBalance >= minLoan{
		num++
		if availableBalance < maxSignleOrderAmount{ // last order
			availableBalance = 0
			orders= append(orders, MussinaOrder{Amount:availableBalance, Rate:0, Period:2})
		}else{
			availableBalance -= maxSignleOrderAmount
			orders= append(orders, MussinaOrder{Amount:maxSignleOrderAmount, Rate:0, Period:2})
		}
	}
	return num
}

