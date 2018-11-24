package main

import (
	"github.com/kr/pretty"
	"github.com/sothonsit/go-tdd/number"
)

func main() {
	a := []int{10, 20, 30, 40}
	n := &number.Num{}
	pretty.Println("Hello word")
	pretty.Println(n.Add(1, 2))
	pretty.Println(n.AddSubstract(1, 2))

	for k, v := range a {
		pretty.Println(k, v)
	}

}
