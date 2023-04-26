package main

import "fmt"

func main() {
	forRangeKeyValue()
	forRangeIgnoreKey()
	forRangeMapKey()
	forRangeIsACopy()
}

func forRangeKeyValue() {
	fmt.Println("for-range loop, print index and value")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}

func forRangeIgnoreKey() {
	fmt.Println("for-range loop, print value only")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func forRangeMapKey() {
	fmt.Println("for-range loop over map, print key only")
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}

func forRangeIsACopy() {
	fmt.Println("for-range loop, show that modifying value variable doesn't modify the original slice")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
}
