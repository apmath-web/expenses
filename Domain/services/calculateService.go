package services

import (
	"errors"
	"github.com/apmath-web/expenses/Domain"
	"math"
)

type CalculateService struct {
	MaxValue float64
	DecimalPlaces int
}

func (cs *CalculateService) GetMaxValue() float64 {
	return cs.MaxValue
}

func (cs *CalculateService) SetMaxValue(maxValue float64) {
	cs.MaxValue = maxValue
}

func (cs *CalculateService) GetDecimalPlaces() int {
	return cs.DecimalPlaces
}

func (cs *CalculateService) SetDecimalPlaces(dp int) {
	cs.DecimalPlaces = dp
}

func (cs *CalculateService) Rounder(maxValue float64) {
	var rounder float64

	pow := math.Pow(10, float64(cs.GetDecimalPlaces()))
	intermed := maxValue * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	tmp := rounder / pow
	cs.SetMaxValue(tmp)
}

func (cs *CalculateService) Calculate(persons []Domain.PersonDomainModelInterface) error{
	var maxValue float64
	var sumWagePerson int

	cs.SetDecimalPlaces(2)

	for _, value := range persons {
		if value.GetSumWage() <= 0 {
			return errors.New("Wage is negative value or zero")
		}

		sumWagePerson += value.GetSumWage()
	}

	maxValue = float64(sumWagePerson / len(persons))
	cs.Rounder(maxValue)

	return nil
}