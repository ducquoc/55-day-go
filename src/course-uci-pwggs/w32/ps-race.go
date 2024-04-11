// Course 3 week 2
package main

import (
	"fmt"
	"time"
)

// RaceCondition - shared variable count is "increased" but not atomic, possible read/write issuer when multi threads
func RaceCondition(count *int) {
	*count++
	fmt.Println(*count)
}

// If you run this program, it will not print 1 2 3 4 5 always, it's likely different orders - demo of race conditions
func main() {

	cnt := 0
	go RaceCondition(&cnt)
	go RaceCondition(&cnt)
	go RaceCondition(&cnt)
	go RaceCondition(&cnt)
	go RaceCondition(&cnt)
	time.Sleep(time.Second)
}
