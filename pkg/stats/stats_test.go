package stats

import (
	"github.com/DariaYudina004/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
	}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "food"},
	}
	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
		{ID: 1, Category: "food"},
	}
	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 2, Category: "food", Amount: 20},
		{ID: 3, Category: "auto", Amount: 30},
		{ID: 4, Category: "auto", Amount: 40},
		{ID: 5, Category: "fun", Amount: 50},
	}
	expected := map[types.Category]types.Money{
		"auto": 80,
		"food": 20,
		"fun":  50,
	}
	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result , expected: %v, actual :%v", expected, result)
	}

}

func TestCategoriesAvg(t *testing.T) {

	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 10},
		{ID: 2, Category: "food", Amount: 20},
		{ID: 3, Category: "auto", Amount: 30},
		{ID: 4, Category: "auto", Amount: 50},
		{ID: 5, Category: "fun", Amount: 50},
		{ID: 5, Category: "fun", Amount: 50},
		{ID: 2, Category: "food", Amount: 40},
	}
	expected := map[types.Category]types.Money{
		"auto": 30,
		"food": 30,
		"fun":  50,
	}
	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result , expected: %v, actual :%v", expected, result)
	}

}

func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	second := map[types.Category]types.Money{
		"auto":   5,
		"food":   3,
		"mobile": 10,
	}

	expected := map[types.Category]types.Money{
		"auto":   -5,
		"food":   -17,
		"mobile": 10,
	}
	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result , expected: %v, actual :%v", expected, result)
	}

}
