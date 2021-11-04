package stats

import "github.com/DariaYudina004/bank/pkg/types"

//рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Amount > 0 {
			sum += payment.Amount
		}
	}
	sum = sum / types.Money(len(payments))
	return sum
}

//сумма покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
