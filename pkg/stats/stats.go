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

//FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {

		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered

}

//CategoriesTotal возвращает сумму платяжей по каждой категории
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount

	}
	return categories

}

//func CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	number := map[types.Category]types.Money{}
	categories := map[types.Category]types.Money{}
	for _, payment := range payments {

		categories[payment.Category] += payment.Amount
		number[payment.Category]++ //len(payments) //считаются только только те,чей статус не Statusfail -->number ++ замена len(payments)
	}

	for key := range categories {
		categories[key] /= number[key]
	}

	return categories
}

//func PeriodsDynamic сравнивает расхлдыы покатегориям за два периода
func PeriodsDynamic(
	first map[types.Category]types.Money, second map[types.Category]types.Money,
) map[types.Category]types.Money {
	amount := map[types.Category]types.Money{}

	for sum := range second {
		amount[sum] += second[sum]
	}

	for sum := range first {
		amount[sum] -= first[sum]
	}
	return amount

}
