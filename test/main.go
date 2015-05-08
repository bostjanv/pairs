package main

import "fmt"
import "log"
import "github.com/bostjanv/pairs/pairs"

func main() {
    a := pairs.MakeSortedArray(10, 6)
    for i := 0; i < len(a); i++ {
        a[i]++
    }
    fmt.Println("Array:", a)

	ps1 := pairs.Find(a, 7)
    fmt.Println("Find:", ps1)

	ps2 := pairs.FindBruteForce(a, 7)
    fmt.Println("FindBruteForce:", ps1)

	if pairs.CheckEqual(ps1, ps1) != true {
		log.Fatal("ps1 != ps2")
	}

	ps3 := pairs.FindUnique(a, 7)
    fmt.Println("FindUnique:", ps3)

	ps2 = pairs.FilterDup(ps2)

	if pairs.CheckEqual(ps2, ps3) != true {
		log.Fatal("ps2 != ps3")
	}
}
