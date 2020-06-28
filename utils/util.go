package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"
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

func CnvTimestamp2String(timestamp int64) string{
	timestamp = timestamp / 1000
	t := time.Unix(timestamp, 0)
	strDate := t.Format(time.RFC3339)
	return strDate
}
