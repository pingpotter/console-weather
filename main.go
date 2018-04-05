package weather

import (
	"strconv"
)

var digitalTop = [10]string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
var digitalMid = [10]string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
var digitalBot = [10]string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}

func toDigital(number int) string {

	str := strconv.Itoa(number)

	return retStr(str, digitalTop) + "\n" + retStr(str, digitalMid) + "\n" + retStr(str, digitalBot)
}

func retStr(str string, digital [10]string) string {

	var ret string

	for _, strN := range str {
		numberFromStr, _ := strconv.Atoi(string(strN))
		ret = ret + digital[numberFromStr]
	}

	return ret
}
