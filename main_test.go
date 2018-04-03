package weather

import "testing"

func Test9Print9(t *testing.T) {

	digital := toDigital(9)
	expectedDigital := " _ \n" +
		"|_|\n" +
		" _|"

	if digital != expectedDigital {
		t.Error("expected \n" + expectedDigital + " actual is " + digital)
	}
}

func Test8Print8(t *testing.T) {

	digital := toDigital(8)
	expectedDigital := " _ \n" +
		"|_|\n" +
		"|_|"

	if digital != expectedDigital {
		t.Error("expected \n" + expectedDigital + " actual is " + digital)
	}
}

func Test89Print89(t *testing.T) {

	digital := toDigital(89)
	expectedDigital := " _  _ \n" +
		"|_||_|\n" +
		"|_| _|"

	if digital != expectedDigital {
		t.Error("expected \n" + expectedDigital + " actual is " + digital)
	}
}
