package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func CnvStruct2Json(v interface{}) string{
	result, err := json.Marshal(v)
	if err != nil {
		log.Println("[ERROR] GetGeneralResp error...", err)
		return fmt.Sprintf("Conv Obj to json fail...%s", err.Error())
	}
	return string(result)
}
