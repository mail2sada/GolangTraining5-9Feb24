package main

import (
	"fmt"
	"training/greet"
	"training/math"
	"training/math/adv"

	"github.com/HuKeping/rbtree"
	"github.com/mail2sada/testpublish/pub"
)

func init() {
	fmt.Println("Main Package init")
}

func main() {
	fmt.Println("I am in main package")
	sum := math.AddMany(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)

	diff := math.Subtract(10, 5)
	testfortrainig()

	diff = math.Sub(10, 5)

	squre := adv.Squere(10)
	fmt.Println(squre)

	fmt.Println(diff)
	greet.Greet("Anand")
	pub.TestPublishMethod("hi")
}

func WorkWithTree() {
	rbt := rbtree.New()
	rbt.Insert(rbtree.Int(10))
	rbt.Insert(rbtree.Int(50))
}
