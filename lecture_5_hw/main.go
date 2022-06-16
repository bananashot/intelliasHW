package main

import "fmt"

var (
	applePrice float64 = 5.99
	pearPrice  float64 = 7
	budget     float64 = 23
)

func main() {

	var (
		appleCount uint = 9
		pearCount  uint = 8
	)

	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	moneyRequired := (applePrice * float64(appleCount)) + (pearPrice * float64(pearCount))
	fmt.Println("-->", moneyRequired)

	fmt.Println("==============")

	fmt.Println("2. Скільки груш ми можемо купити?")
	fmt.Println("-->", int(budget/pearPrice))

	fmt.Println("==============")

	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Println("-->", int(budget/applePrice))

	fmt.Println("==============")

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	appleCount = 2
	pearCount = 2
	moneyCheck := (applePrice * float64(appleCount)) + (pearPrice * float64(pearCount))
	fmt.Println("-->", budget >= moneyCheck)

}
