package main

import (
	//"sync"
	"time"
)

var (
	money = 100
	//lock  = sync.Mutex{}
)

func spendyOne() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money -= 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("SpendyOne Done")
}

func stingyOne() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money += 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("StingyOne Done")
}

func spendyTwo() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money -= 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("SpendyTwo Done")
}

func stingyTwo() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money += 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("StingyTwo Done")
}

func spendyThree() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money -= 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("SpendyThree Done")
}

func stingyThree() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money += 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("StingyThree Done")
}

func spendyFour() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money -= 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("SpendyFour Done")
}

func stingyFour() {
	for i := 1; i <= 1000; i++ {
		//lock.Lock()
		money += 10
		//lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("StingyFour Done")
}

// func spendyFive() {
// 	for i := 1; i <= 1000; i++ {
// 		//lock.Lock()
// 		money -= 10
// 		//lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("SpendyFive Done")
// }

// func stingyFive() {
// 	for i := 1; i <= 1000; i++ {
// 		//lock.Lock()
// 		money += 10
// 		//lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("StingyFive Done")
// }

// func spendySix() {
// 	for i := 1; i <= 1000; i++ {
// 		//lock.Lock()
// 		money -= 10
// 		//lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("SpendySix Done")
// }

// func stingySix() {
// 	for i := 1; i <= 1000; i++ {
// 		//lock.Lock()
// 		money += 10
// 		//lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("StingySix Done")
// }

// func spendySeven() {
// 	for i := 1; i <= 1000; i++ {
// 		//lock.Lock()
// 		money -= 10
// 		//lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("SpendySeven Done")
// }

// func stingySeven() {
// 	for i := 1; i <= 1000; i++ {
// 		//lock.Lock()
// 		money += 10
// 		//lock.Unlock()
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("StingySeven Done")
// }

func main() {
	go spendyOne()
	go stingyOne()

	go spendyTwo()
	go stingyTwo()

	go spendyThree()
	go stingyThree()

	go spendyFour()
	go stingyFour()

	// go spendyFive()
	// go stingyFive()

	// go spendySix()
	// go stingySix()

	// go spendySeven()
	// go stingySeven()

	time.Sleep(3000 * time.Millisecond)
	print("total money : ", money)
}
