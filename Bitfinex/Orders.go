package Bitfinex

import "math"

type MussinaOrder struct{
	Amount float64
	Rate float64
}



func GenOrders(availableBalance float64, maxSignleOrderAmount float64, left float64) int{
	availableBalance -= left
	orders := make([]MussinaOrder, 0)
	num := 0
	for availableBalance > 0{
		num++
		if availableBalance < maxSignleOrderAmount{ // last order
			availableBalance = 0
			orders= append(orders, MussinaOrder{Amount:availableBalance, Rate:0})
		}else{
			availableBalance -= maxSignleOrderAmount
			orders= append(orders, MussinaOrder{Amount:maxSignleOrderAmount, Rate:0})
		}
	}
	return num
}

func calNumOfOrders(availableBalance float64, maxSignleOrderAmount float64) int64{
	num := int64(availableBalance / maxSignleOrderAmount)

	if math.Mod(availableBalance, maxSignleOrderAmount) != 0{
		num += 1
	}
	return num
}
