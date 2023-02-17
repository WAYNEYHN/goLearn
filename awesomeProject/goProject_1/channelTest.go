package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	myChan := make(chan int, 2)
	myChan <- 1
	close(myChan)
	_, ok := <-myChan
	fmt.Println(ok)
	testSelect()

}

func testSelect() {
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}

func recevingChanWithSelect() {
	chanSlice := []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	randInt := rand.Intn(5)
	if randInt < 2 {
		chanSlice[randInt] <- 1
	}

	select {
	case <-chanSlice[0]:
		fmt.Println("choice: 0")
	case <-chanSlice[1]:
		fmt.Println("choice 1")
	case <-chanSlice[2]:
		fmt.Println("choice 2")
	default:
		fmt.Println("choice default")
	}
}

func singleChanelTest(a chan<- int) {
	a <- 1
	//fmt.Println(m)
}

// 判断管道的数据传输时深拷贝还是浅拷贝
func checkPassValue() {
	a := []int{1, 2}
	b := []int{2, 3}
	var c1 = make([]int, 3)
	copy(c1, a)
	var secondChan = make(chan []int, 2)
	fmt.Printf("%p, %p, %p\n", &a, &b, a)
	secondChan <- a
	secondChan <- b
	c := <-secondChan
	d := <-secondChan
	fmt.Printf("%p, %p, %p\n", &c, &d, c1)
	fmt.Println(c1)
}

func testChannal_1() {
	firstChan := make(chan int, 6)
	chanInit(firstChan)
	//close(firstChan)
	go func(firstChan chan int) {
		time.Sleep(time.Second * 5)
		firstChan <- 1
	}(firstChan)
	traverseChan(firstChan) //遍历通道中剩余元素
}

// 初始化管道， 生产比管道容量少一位的元素
func chanInit(firstChan chan int) {
	for i := 0; i < cap(firstChan)-1; i++ {
		firstChan <- i
	}
}

// 消费者，遍历管道元素
/*
	firstChan 管道
	num 需要消费的管道中元素
*/
func consumer(firstChan chan int, num int) error {
	fmt.Printf("cap of chan: %d\nlen of chan: %d\n", cap(firstChan), len(firstChan))
	for i := 0; i < num; i++ {
		tmp, ok := <-firstChan
		if ok {
			fmt.Printf("element in chan: %d isEmpty: %t\n", tmp, ok)
		} else {
			fmt.Printf("element in chan: %d isEmpty: %t\n", tmp, ok)
			return errors.New("chan is empty")
		}

	}
	return nil
}

// for-range 遍历管道
func traverseChan(a chan int) {
	for c := range a {
		fmt.Printf("last elements in chan: %d\n", c)
		//fmt.Printf("Pointer of c: %p\n Pointer of a: %p\n", &c, &a)
	}
}

// 未初始化的管道进行发送操作会阻塞
func testBlock() {
	var secondChan chan int
	secondChan <- 1
}
