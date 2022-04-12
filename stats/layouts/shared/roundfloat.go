package shared

import (
	"fmt"
	"strconv"
)

func RoundFloat(flt string) interface{} {
	float, _ := strconv.ParseFloat(flt, 64)
	return fmt.Sprintf("%.2f", float)
}

func FloatToInt(flt string) interface{} {
	float, _ := strconv.ParseFloat(flt, 64)
	return fmt.Sprint(int(float))
}
