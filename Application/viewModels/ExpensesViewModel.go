package viewModels

import (
	"encoding/json"
	"github.com/apmath-web/expenses/Application/Validation"
	"github.com/apmath-web/expenses/Domain"
)

type JsonExpenses struct {
	amount float64 `json:"maxPayment"`
}

type ExpensesViewModel struct {
	JsonExpenses
	validation Validation.Validation
}

func (expensesViewModel *ExpensesViewModel) GetAmount() float64 {
	return expensesViewModel.amount
}

func (expensesViewModel *ExpensesViewModel) validateAmount() bool {
	if expensesViewModel.amount < 0 {
		expensesViewModel.validation.AddMessage(Validation.GenMessage("amount", "Is negative"))
		return false
	}
	return true
}

func (expensesViewModel *ExpensesViewModel) Validate() bool {
	if expensesViewModel.validateAmount() {
		return true
	}
	return false
}

func (expensesViewModel *ExpensesViewModel) MarshalJSON() (b []byte, e error) {
	return json.Marshal(map[string]interface{}{
		"maxPayment": expensesViewModel.amount,
	})
}

func (expensesViewModel *ExpensesViewModel) Hydrate(expensesInerface Domain.ExpensesInterface) {
	expensesViewModel.amount = expensesInerface.GetAmount()
}
