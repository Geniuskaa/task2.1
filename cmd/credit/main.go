package main

import (
	"fmt"
	"github.com/Geniuskaa/task2.1/pkg/credit"
)

func main() {
	monthlyPayment, interestPayment, totalPayment := credit.Calculate(1_000_000_00, 20, 3)
	fmt.Println(monthlyPayment)
	fmt.Println(interestPayment)
	fmt.Println(totalPayment)

}
