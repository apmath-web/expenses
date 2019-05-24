package services

import "testing"

func TestCalculateService_Rounder(t *testing.T) {
	var x float64
	var per int
	var calc = CalculateService{}
	var answer float64

	x = 12.3456
	per = 2
	calc.SetDecimalPlaces(per)
	calc.Rounder(x)
	answer = calc.GetMaxValue()

	if answer != 12.35 {
		t.Error("Incorrect behavior with positive value")
	}

	x = -7.87653
	per = 3
	calc.SetDecimalPlaces(per)
	calc.Rounder(x)
	answer = calc.GetMaxValue()

	if answer != -7.877 {
		t.Error("Incorrect behavior with negative value")
	}

}
