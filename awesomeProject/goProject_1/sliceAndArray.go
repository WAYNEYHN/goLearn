package main

import (
	"fmt"
	"reflect"
)

var (
	a int
)

/*
 */
func sliceStatic() {
	var slice1 = []string{"123", "123"}
	var slice2 = []int{6: 1}
	a := reflect.TypeOf(slice1)
	fmt.Println(a, slice1, slice2)
}

/*
扩容机制验证
*/
func sliceTest() {
	var slice1 = make([]uint16, 4)
	fmt.Printf("%d, %d, %p\n", len(slice1), cap(slice1), slice1)
	slice1[0] = 1
	for i := 0; i < 10; i++ {
		oldCap := cap(slice1)
		slice1 = append(slice1, 1)
		fmt.Printf("%f, %d, %p, %p\n", float32(cap(slice1))/float32(oldCap), cap(slice1), slice1, &slice1)
	}
}

func sliceExpand() {
	var slice1 = make([]uint16, 1024)
	fmt.Printf("%d, %d, %p\n", len(slice1), cap(slice1), slice1)
	for i := 0; i < 10; i++ {
		oldCap := cap(slice1)
		slice1 = append(slice1, slice1[:len(slice1)/4]...)
		fmt.Printf("%f, %d, %p, %p\n", float32(cap(slice1))/float32(oldCap), cap(slice1), slice1, &slice1)
	}
}

func newSlice() {
	var slice1 = make([]int, 2, 2)
	var slice2 = slice1
	fmt.Printf("%p, %p\n", &slice1, &slice2)
	slice2 = append(slice2, slice2[:]...)
	fmt.Println(slice1, slice2)
	fmt.Printf("%p, %p\n", &slice1, &slice2)

	var array1 = [3]int{1, 3, 6}
	var arrayPointer = &array1
	fmt.Println(*arrayPointer)

}

func arrayTest() {
	var array1 = [3]int{1, 3, 6}
	var array2 = [...]int{6: 1}
	var array3 = array1[0:2]
	fmt.Println(array1, array2, array3)
	a := reflect.TypeOf(array3)
	fmt.Println(a, a.Name(), a.Kind(), a.Elem())
}

func simeple() {
	var int_1 = 123
	int_2 := "123"
	var int_3 int
	fmt.Println(int_1, int_2, int_3)
}
