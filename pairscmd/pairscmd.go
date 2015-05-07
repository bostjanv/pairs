package main

import "fmt"
import "log"
import "github.com/bostjanv/pairs/pairs"

func main() {
	fmt.Println("Find all pairs in a sorted array that sums to a given value")
	
	a := pairs.MakeSortedArray(10000, 10000)
	//fmt.Println(a)
	
	ps1 := pairs.Find(a, 100)
	//fmt.Println(ps)
	
	ps2 := pairs.FindBruteForce(a, 100)
	//fmt.Println(ps1)

	if pairs.CheckEqual(ps1, ps1) != true {
		log.Fatal("ps1 != ps2")
	}

	ps3 := pairs.FindUnique(a, 100)
	ps2 = pairs.FilterDup(ps2)

	if pairs.CheckEqual(ps2, ps3) != true {
		log.Fatal("ps2 != ps3")
	}
}
