package Bitfinex

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"log"
	"time"
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

func ModifyPeriod(orders *[]MussinaOrder, loan30DaysRate float64) *[]MussinaOrder{
	log.Println("[ModifyPeriod]...")
	for index, _ := range *orders{
		rate := (*orders)[index].Rate
		annualRate := rate * 365
		if annualRate > loan30DaysRate{
			(*orders)[index].Period = 30
			log.Println(fmt.Sprintf("[ModifyPeriod] amount:[%f] rate:[%f] change period to 30.",
				(*orders)[index].Amount, (*orders)[index].Rate))
		}
	}
	return orders
}

func SubmitOrders(orders *[]MussinaOrder) *[]bitfinex.Notification{
	log.Println(fmt.Sprintf("[SubmitOrders] orders size:[%d]", len(*orders)))
	bfOffersNoti := make([]bitfinex.Notification, 0, 10)
	for _, order := range *orders{
		noti,err := bfRestClient.Funding.SubmitOffer(&bitfinex.FundingOfferRequest{
			Type:"LIMIT",
			Symbol:"fUSD",
			Amount:order.Amount,
			Rate:order.Rate * 0.01,  // bitfinex rate is raw data. ex: 0.036 % ---> 0.00036
			Period:order.Period,
			Hidden:false,
		})
		if err != nil{
			log.Println("[ERROR] ", err.Error())
		}else{
			log.Println("[SubmitOrders] succ...")
			bfOffersNoti = append(bfOffersNoti, *noti)
			log.Println(noti)
		}
	}
	return &bfOffersNoti
}

func GetActiveOrdersSize() int{
	log.Println("[GetActiveOrdersSize]...")
	snapshot, err := bfRestClient.Funding.Offers("fUSD")
	if err != nil{
		log.Println("[ERROR] ", err.Error())
		return 0
	}else{
		if snapshot != nil{
			for _, offer := range snapshot.Snapshot{
				log.Println(*offer)
			}
			log.Println(fmt.Sprintf("active order size:[%d]", len(snapshot.Snapshot)))
			return len(snapshot.Snapshot)
		}else{
			log.Println("No active orders...")
			return 0
		}
	}
}

func GetActiveOrders() *[]*bitfinex.Offer{
	log.Println("[GetAllActiveOrders]...")
	snapshot, err := bfRestClient.Funding.Offers("fUSD")
	if err != nil{
		log.Println("[ERROR] ", err.Error())
		return nil
	}else{
		if snapshot != nil{
			log.Println(fmt.Sprintf("active order size:[%d]", len(snapshot.Snapshot)))
			return &snapshot.Snapshot
		}else{
			log.Println("No active orders...")
			return nil
		}
	}
}

func CancelAllOrders(orders *[]*bitfinex.Offer) {
	log.Println("[CancelAllOrders]...orders size:", len(*orders))
	for _, order := range *orders{
		resp, err := bfRestClient.Funding.CancelOffer(&bitfinex.FundingOfferCancelRequest{
			Id:order.ID,
		})
		if err != nil{
			log.Println("[ERROR] ", err.Error())
		}else{
			log.Println(resp.Text)
		}
		time.Sleep(100 * time.Millisecond)
	}

}

func GetActivePositions() *bitfinex.FundingCreditSnapshot{
	log.Println("[GetActivePositions]...")
	credit, err := bfRestClient.Funding.Credits("fUSD")
	if err != nil{
		log.Println("[ERROR] GetActivePositions...", err.Error())
	}
	return credit
}

func GetLedgers(currency string, start int64, end int64) *[]*bitfinex.Ledger{
	log.Println("[GetLedgers]...", currency)
	ledgersSnap, err := bfRestClient.Ledgers.Ledgers(currency, start, end, 500)
	if err != nil{
		log.Println("[ERROR] GetLedgers...", err.Error())
		return nil
	}
	if ledgersSnap == nil{
		log.Println("[ERROR] GetLedgers...ledger snapshot is nil")
		return nil
	}
	log.Println("ledger size:", len(ledgersSnap.Snapshot))
	return &ledgersSnap.Snapshot
}