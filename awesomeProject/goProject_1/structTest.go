package main

import "fmt"

func main() {
	stuctTest()
}

type fruitNum struct {
	appleNum  int
	bananaNum int
}

func (a fruitNum) String() string {
	return fmt.Sprintf("appale: %d, banana: %d\n", a.appleNum, a.bananaNum)

}

func (a fruitNum) myAdd() int {
	return a.bananaNum + a.appleNum
}

type fruits struct {
	fruitsName []string
	fruitNum
}

func (a *fruits) myInit(names []string) {
	a.fruitsName = names
}

func (a fruits) String() string {
	return fmt.Sprintf("fruitsName: %s\n  fruitsNum: %d, %d\n", a.fruitsName, a.appleNum, a.bananaNum)
}

func stuctTest() {

	typeReal := fruitNum{
		appleNum:  1,
		bananaNum: 2,
	}
	fmt.Printf("the num of fruits: %s\n", typeReal)
	fmt.Printf("total num of fruits: %d\n", typeReal.myAdd())

	fruitName := []string{"origin", "bee"}
	fruitName2 := []string{"origins", "bees"}
	buyFruits := fruits{
		fruitsName: fruitName,
		fruitNum:   typeReal,
	}

	buyFruits.myInit(fruitName2)
	fmt.Printf("%s", buyFruits)
	fmt.Printf("%s", buyFruits.fruitNum.String())
}
