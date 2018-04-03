package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak many"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
}

func weatherCelsius(celsius int, desc string) string {

	strCel := strconv.Itoa(celsius)

	var str [][]string

	for _, point := range strCel {
		str = append(str, pattern(string(point)))
	}

	str = append(str, pattern("c"))

	ret := printOut(str)
	ret += desc

	return ret
}

func printOut(digit [][]string) string {

	ret := ""
	for n := 0; n < 3; n++ {
		for m := 0; m < len(digit); m++ {
			ret += digit[m][n]
		}
		ret += "\n"
	}
	return ret
}
func pattern(number string) []string {

	ret := []string{}

	switch number {
	case "-":
		ret = []string{" ", "_", " "}
	case ".":
		ret = []string{" ", " ", "."}
	case "c":
		ret = []string{"   ", "   ", " C "}
	case "1":
		ret = []string{"   ", "  |", "  |"}
	case "2":
		ret = []string{" _ ", " _|", "|_ "}
	case "3":
		ret = []string{" _ ", " _|", " _|"}
	case "4":
		ret = []string{"   ", "|_|", "  |"}
	case "5":
		ret = []string{" _ ", "|_ ", " _|"}
	case "6":
		ret = []string{" _ ", "|_ ", "|_|"}
	case "7":
		ret = []string{" _ ", "  |", "  |"}
	case "8":
		ret = []string{" _ ", "|_|", "|_|"}
	case "9":
		ret = []string{" _ ", "|_|", " _|"}
	default:
		ret = []string{" _ ", "| |", "|_|"}
	}
	return ret
}
