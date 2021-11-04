package stats

import (
	"fmt"

	"github.com/DariaYudina004/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 1000_00,
		},
		{
			Amount: 2000_00,
		},
		{
			Amount: 2000_00,
		},
		{
			Amount: 3000_00,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	//200000

}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Amount:   1000_00,
			Category: "auto",
		},
		{
			Amount:   2000_00,
			Category: "auto",
		},
		{
			Amount:   2000_00,
			Category: "auto",
		},
		{
			Amount:   1000_00,
			Category: "store",
		},
	}
	fmt.Println(TotalInCategory(payments,types.Category("auto")))
	//Output:
	//500000

}
