package stats

import (
	"fmt"
	"github.com/DariaYudina004/bank/v2/pkg/types"
	"testing"
	"reflect"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Status: types.StatusFail,
			Amount: 1000_00,
		},
		{
			Status: types.StatusFail,
			Amount: 2000_00,
		},
		{
			Status: types.StatusInProgress,
			Amount: 2000_00,
		},
		{
			Status: types.StatusOk,
			Amount: 4000_00,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	//300000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Status:   types.StatusFail,
			Amount:   1000_00,
			Category: "auto",
		},
		{
			Status:   types.StatusInProgress,
			Amount:   2000_00,
			Category: "auto",
		},
		{
			Status:   types.StatusOk,
			Amount:   2000_00,
			Category: "auto",
		},
		{
			Status:   types.StatusOk,
			Amount:   1000_00,
			Category: "store",
		},
	}
	fmt.Println(TotalInCategory(payments, types.Category("auto")))
	//Output:
	//400000

}


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
