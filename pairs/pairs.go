package pairs

import "fmt"
import "math/rand"
import "sort"

// Find unique pairs in array that sum to s.
func FindUnique(a []int, s int) (ps [][2]int) {
	ps = make([][2]int, 0)
	i, j := 0, len(a)-1

	c := 0 // loop iteration counter

	for j > i {
		//fmt.Printf("Comparing %d and %d\n", a[i], a[j])
		c++
		s1 := a[i] + a[j]
		if s1 == s {
			p := [2]int{a[i],a[j]}
			ps = append(ps, p)
			ai := a[i]
			// Skip all elements that are equal to a[i].
			for i++; i < j && a[i] == ai; i++ {
			}
			j--
		} else if s1 > s {
			j--
		} else if s1 < s {
			i++
		}
	}

	fmt.Println("[FindUnique]", c, len(ps))

	return
}

// Finds all pairs (duplicates included) in array that sum to s.
func Find(a []int, s int) (ps [][2]int) {
	ps = make([][2]int, 0)
	i, j := 0, len(a)-1

	ndup := 0 // duplicates counter
	c := 0 // iteration counter

	for j > i {
		//fmt.Printf("Comparing %d and %d\n", a[i], a[j])
		c++
		s1 := a[i] + a[j]
		if s1 == s {
			p := [2]int{a[i],a[j]}
			ps = append(ps, p)
			ndup++
			i++
		} else if s1 > s {
			i = i-ndup
			ndup = 0
			j--
		} else if s1 < s {
			i++
		}
	}

	fmt.Println("[Find]", c, len(ps))

	return
}

func FindBruteForce(a []int, s int) (ps [][2]int) {
	ps = make([][2]int, 0)

	c := 0

	for i := 0; i < len(a); i++ {
		for j := i+1; j < len(a); j++ {
			c++
			s1 := a[i] + a[j]
			if s1 == s {
				p := [2]int{a[i],a[j]}
				ps = append(ps, p)
			} else if s1 > s  {
				break
			}
		}
	}
	
	fmt.Println("[FindBruteForce]", c, len(ps))

	return
}

func CheckEqual(a, b [][2]int) bool {
	n := len(a)
	if n != len(b) {
		return false
	}

	for i := 0; i < n; i++ {
		if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
			return false
		}
	}

	return true
}

func FilterDup(a [][2]int) (b [][2]int) {
	n := len(a)
	b = make([][2]int, 0, n)

	b = append(b, a[0])
	for i := 1; i < n; i++ {
		if a[i][0] != a[i-1][0] || a[i][1] != a[i-1][1] {
			b = append(b, a[i])
		}
	}

	return
}

// Creates a sorted array of ints in increasing order.
// Array is allowed to contain duplicates.
func MakeSortedArray(n, m int) (a []int) {
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(m) 
	}
	
	// We don't bother to implement our own version of sort.
	sort.Ints(a)
	
	return
}