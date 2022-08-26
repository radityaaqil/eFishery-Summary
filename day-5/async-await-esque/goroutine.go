package main

import (
	"fmt"
	"sync"
)

//Race Condition Solutions
	var saldo int = 0

	//Solution 1
	var isLock bool = false

	//Solution 2
	var lock sync.Mutex

	func addSaldo(){
		for i := 0; i < 1000000; i++ {
			saldo++
		}
	}

	//Solution 3

	// func addSaldo(){
	// 	for i := 0; i < 1000000; i++ {
	// 		atomic.AddInt64(&saldo, 1)
	// 	}
	// }

	func Lock(lockedFunc func()) {
		for isLock {
			if !isLock {
				break
			}
		}
		isLock = true
		lockedFunc()
		isLock = false
	}

func main() {
	//WG similar to async await in nodeJS

	// wg := sync.WaitGroup{}

	// wg.Add(2)

	// go func(){
	// 	defer wg.Done()
	// 	fmt.Println("Hello")
	// }()

	// go func(){
	// 	defer wg.Done()
	// 	fmt.Println("World")
	// }()

	// go func(){
	// 	defer wg.Done()
	// 	fmt.Println("Dang")
	// }()

	// wg.Wait()

	//===============================================================================================================

	//Race Condition --> anomaly caused by goroutine (multiple simultaneous processes) when accessing same variable
	//Solution : Optimistic Locking & Pessimistic Locking

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		// Solution 1
		// Lock(addSaldo)

		//Solution 2
		lock.Lock()
		addSaldo()
		lock.Unlock()

		//Solution 3
		//addSaldo() --> with atomic
	}()

	go func() {
		defer wg.Done()
		// Solution 1
		// Lock(addSaldo)

		//Solution 2
		lock.Lock()
		addSaldo()
		lock.Unlock()

		//Solution 3
		//addSaldo() --> with atomic
	}()

	wg.Wait()
	fmt.Println(saldo)
}

	