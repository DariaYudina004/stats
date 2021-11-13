package stats

import (
	"reflect"
	"testing"

	"github.com/DariaYudina004/bank/v2/pkg/types"
)

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
