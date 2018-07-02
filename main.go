package main

import "fmt"

func main() {
	bob := Cats{"NIK", 7, 0.8988}
	fmt.Println("Bob is", bob, "\n ", "Bob func is ", bob.test())
	fmt.Println(bob)
}

type Cats struct {
	name      string
	age       int
	happiness float64
}

func (cat *Cats) test() float64 {
	cat.age = 123
	return float64(cat.age) * cat.happiness
}
