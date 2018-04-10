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

func toDigitalPing(number int) string {

	str := strconv.Itoa(number)

	return retStr(str, digitalTop) + "\n" + retStr(str, digitalMid) + "\n" + retStr(str, digitalBot)
}

var fourtyeight = 48

func toDigitalPack(number int) string {
	str := strconv.Itoa(number)
	palldigit := finddigit(str)
	return palldigit
}

func finddigit(str string) string {
	var digit int
	var alltop string
	var allmid string
	var allbot string

	for i := range str {
		digit = int(str[i])
		allmid += digitalMid[digit-fourtyeight]
		alltop += digitalTop[digit-fourtyeight]
		allbot += digitalBot[digit-fourtyeight]
	}
	alldigit := alltop + "\n" + allmid + "\n" + allbot
	return alldigit
}

func toDigitalBell(number int) string {

	str := strconv.Itoa(number)

	var retIntTop string
	var retIntMid string
	var retIntBtm string

	for _, value := range str {
		toInt, _ := strconv.Atoi(string(value))
		retIntTop += digitalTop[toInt]
		retIntMid += digitalMid[toInt]
		retIntBtm += digitalBot[toInt]
	}

	return retIntTop + "\n" + retIntMid + "\n" + retIntBtm
}
