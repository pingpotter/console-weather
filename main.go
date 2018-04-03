package weather

import (
	"strconv"
)

func toDigital(number int) string {

	digitalTop := [10]string{" _ ", "", "", "", "", "", "", "", " _ ", " _ "}
	digitalMid := [10]string{"| |", "", "", "", "", "", "", "", "|_|", "|_|"}
	digitalBot := [10]string{"|_|", "", "", "", "", "", "", "", "|_|", " _|"}

	str := strconv.Itoa(number)
	var ret string

	for _, strN := range str {
		numberFromStr, _ := strconv.Atoi(string(strN))
		ret = ret + digitalTop[numberFromStr]
	}
	ret = ret + "\n"
	for _, strN := range str {
		numberFromStr, _ := strconv.Atoi(string(strN))
		ret = ret + digitalMid[numberFromStr]
	}
	ret = ret + "\n"
	for _, strN := range str {
		numberFromStr, _ := strconv.Atoi(string(strN))
		ret = ret + digitalBot[numberFromStr]
	}

	return ret
}
