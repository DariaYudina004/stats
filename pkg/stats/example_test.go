package stats

import (
	"fmt"

	"github.com/DariaYudina004/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			// Status : "FAIL",
			Amount: 1000_00,
		},
		{
			// Status : "FAIL",
			Amount: 2000_00,
		},
		{
			// Status : "INPROGRESS",
			Amount: 2000_00,
		},
		{
			// Status : "OK",
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
			// Status : "FAIL",
			Amount:   1000_00,
			Category: "auto",
		},
		{
			// Status : "INPROGRESS",
			Amount:   2000_00,
			Category: "auto",
		},
		{
			// Status : "OK",
			Amount:   2000_00,
			Category: "auto",
		},
		{
			// Status : "OK",
			Amount:   1000_00,
			Category: "store",
		},
	}
	fmt.Println(TotalInCategory(payments,types.Category("auto")))
	//Output:
	//400000

}
