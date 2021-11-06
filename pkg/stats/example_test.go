package stats

import (
	"fmt"
	"github.com/DariaYudina004/bank/v2/pkg/types"
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
