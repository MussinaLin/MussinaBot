package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

func CnvStruct2Json(v interface{}) string{
	result, err := json.Marshal(v)
	if err != nil {
		log.Println("[ERROR] GetGeneralResp error...", err)
		return fmt.Sprintf("Conv Obj to json fail...%s", err.Error())
	}
	return string(result)
}

func CnvDailyRate2AnnualRate(rate float64)float64{
	apy := rate * 365
	return math.Round(apy * 100000) / 100000
}
