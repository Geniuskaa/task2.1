package credit_test

import (
	"fmt"
	"github.com/Geniuskaa/task2.1/pkg/credit"
)

func ExampleCalculateCredit() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.CalculateCredit(1_000_000_00, 20, 3))
	fmt.Println(credit.CalculateCredit(10_000_00, 20, 3))
	// Output:
	// 37164 337904 1337904
	// 372 3392 13392
}

