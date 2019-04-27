package models

import "github.com/apmath-web/expenses/Domain"

type ExpensesDomainModel struct {
	amount float64
}

func (ex *ExpensesDomainModel) GetAmount() float64 {
	return ex.amount
}

func (ex *ExpensesDomainModel) SetAmount(amount float64) {
	ex.amount = amount
}

func GenExpensesDomainModel(amount float64) Domain.ExpensesInterface {
	dm := new(ExpensesDomainModel)
	dm.amount = amount
	return dm
}
