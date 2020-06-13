package main

import (
	"fmt"
	"github.com/Geniuskaa/task2.1/pkg/credit"
)

func main() {
	fmt.Println(credit.CalculateCredit(1_000_000_00, 20, 3))
}
