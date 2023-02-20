package main

import "fmt"

func main() {
	numbers := []int8{1, 2, 3, 4, 5, 6}
	print(numbers)
	//RangeTest()
}

func RangeTest() {
	numbers := [...]int8{1, 2, 3, 4, 5, 6}
	maxIndex := len(numbers) - 1
	for i, nums := range numbers {
		if i == maxIndex {
			numbers[0] += nums
		} else {
			numbers[i+1] += nums
		}

	}
	fmt.Println(numbers)
}

func switchTest() {

	value6 := interface{}(byte(127))
	value6.type()

	switch t := value6.(type) {
	case uint8, uint16:
		fmt.Println("uint8 or uint16")
	case byte:
		fmt.Printf("byte")
	default:
		fmt.Printf("unsupported type: %T", t)
	}
}
