package Bitfinex

import "math"

type MussinaOrder struct{
	Amount float64
	Rate float64
}



func GenOrders(availableBalance float64, maxSignleOrderAmount float64) *[]MussinaOrder{
	numOfOrders := calNumOfOrders(availableBalance, maxSignleOrderAmount)
	orders := make([]MussinaOrder, numOfOrders)
	orders = orders
	return nil
}

func calNumOfOrders(availableBalance float64, maxSignleOrderAmount float64) int64{
	num := int64(availableBalance / maxSignleOrderAmount)

	if math.Mod(availableBalance, maxSignleOrderAmount) != 0{
		num += 1
	}
	return num
}
