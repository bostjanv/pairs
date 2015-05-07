package pairs

import "math/rand"
import "sort"

// Find unique pairs in array that sum to s.
func FindUnique(a []int, s int) (ps [][2]int, c uint) {
	ps = make([][2]int, 0)
	i, j := 0, len(a)-1

	c = 0 // loop iteration counter

	for j > i {
		//fmt.Printf("Comparing %d and %d\n", a[i], a[j])
		c++
		s1 := a[i] + a[j]
		if s1 == s {
			p := [2]int{a[i],a[j]}
			ps = append(ps, p)
			// Skip all elements that are equal to a[i] and a[j]
            ai, aj := a[i], a[j]
			for i++; i < j && a[i] == ai; i++ {
                c++
			}
			for j--; i < j && a[j] == aj; j-- {
                c++
            }
		} else if s1 > s {
			j--
		} else if s1 < s {
			i++
		}
	}

	return
}

// Finds all pairs (duplicates included) in array that sum to s.
func Find(a []int, s int) (ps [][2]int, c uint) {
	ps = make([][2]int, 0)
	i, j := 0, len(a)-1

	c = 0 // iteration counter

	for j > i {
		//fmt.Printf("Comparing %d and %d\n", a[i], a[j])
		s1 := a[i] + a[j]
		if s1 == s {
			p := [2]int{a[i],a[j]}
			ps = append(ps, p)
            ai := a[i]
            ndup := 1
            c++
            for i++; j > i && a[i] == ai; i++ {
			    ps = append(ps, p)
                ndup++
                c++
            }
            i = i-ndup
            j--
		} else if s1 > s {
			j--
		} else if s1 < s {
			i++
		}
	}

	return
}

func FindBruteForce(a []int, s int) (ps [][2]int, c int) {
	ps = make([][2]int, 0)

	c = 0

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

    if n == 0 {
        return
    }

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
