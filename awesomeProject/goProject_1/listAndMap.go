package main

import (
	"container/list"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

//var name1 = "123"

func mapTest() {

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := myMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found")
	}

}

func mapUnThreadSafeTest() {
	testMap := map[string]string{
		"test": "test",
	}
	go read(testMap)
	time.Sleep(time.Second)
	go write(testMap)
	time.Sleep(30 * time.Second)
	fmt.Println(testMap)

}

func read(m map[string]string) {
	for {
		_ = m["test"]
		time.Sleep(1)
	}
}

func write(m map[string]string) {
	for {
		m["test"] = "wrote"
		time.Sleep(1)
	}
}

func mapTraverse() {
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, value := range myMap {
		fmt.Printf("the element of key %q: %d\n", key, value)
	}
	for key := range myMap {
		fmt.Printf("the element of key %q: %d\n", key, myMap[key])
	}
}

func mayError1() {
	var myMap0 map[string]int
	var myMap1 = map[string]int{}
	var myMap2 = make(map[string]int)
	myMap0["123"] = 1
	myMap1["123"] = 1
	myMap2["123"] = 1
	fmt.Println(myMap0, myMap1, myMap2)
}

func demo() {
	var name = flag.String("name", "everyone", "the greeting Object")
	var name1 = "every"
	println(&name1)
	{
		name1 := "everyone"
		println(&name1)
	}
	{
		var name1 = "everyone"
		println(&name1)
	}

	name1, name3 := "everyOne", "every"
	fmt.Println(reflect.TypeOf(name).Kind())
	flag.Parse()
	fmt.Printf("hello, %v, %v\n", *name, &name1, name3)
}

func reState() {
	var err1 error
	n, err1 := io.WriteString(os.Stdout, "hello, everyone\n") //变量重命名, 实际上只是一次赋值操作
	//n, err1 := io.WriteString(os.Stdout, "hello, everyone") //编译error： no new variables on left side of :=
	fmt.Printf("%p\n", &err1)
	fmt.Printf("%d, %p\n", n, &err1)
}

func reflectTest() {
	var list_1 list.List
	string1 := "123"
	var container = [3]string{"zero", "one", "two"}
	listType := reflect.TypeOf(list_1)
	fmt.Println(list_1, listType.Kind(), listType.Name())
	value, ok := interface{}(container).([]string)    //类型断言
	value1, ok1 := interface{}(container).([3]string) //类型断言表达式
	_, ok2 := interface{}(string1).([]string)
	fmt.Println(value, ok)
	fmt.Println(value1, ok1)
	fmt.Println(string1, ok2)

}

func MyReflect(a interface{}) {
	theType := reflect.TypeOf(a)
	fmt.Println(theType.Kind())
}
