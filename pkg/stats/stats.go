package stats

import "github.com/DariaYudina004/bank/v2/pkg/types"

//рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	var number types.Money
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += payment.Amount
		number++ //len(payments) //считаются только только те,чей статус не Statusfail -->number ++ замена len(payments)
	}
	sum = sum / number
	return sum

	//код до появления статуса
	// var sum types.Money
	// for _, payment := range payments {
	// 	sum += payment.Amount
	// }
	// sum = sum / types.Money(len(payments))
	// return sum
}

//сумма покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}

		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}

//CategoriesTotal возвращает сумму платяжей по каждой категории 
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
		
	}
	return categories
}
