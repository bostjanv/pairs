package pairs

import "math/rand"
import "sort"

// Find unique pairs of integers in array that sum to s.
// Time complexity: O(n)
func FindUnique(a []int, s int) (ps [][2]int) {
	ps = make([][2]int, 0)
	i, j := 0, len(a)-1

	for j > i {
		ai, aj := a[i], a[j]
		s1 := ai + aj
		if s1 == s {
			p := [2]int{ai, aj}
			ps = append(ps, p)
			// Skip all elements that are equal to ai and aj
			for i++; i < j && a[i] == ai; i++ {
			}
			for j--; i < j && a[j] == aj; j-- {
			}
		} else if s1 > s {
			j--
		} else if s1 < s {
			i++
		}
	}

	return
}

// Finds all pairs of integers (duplicates included) in array that sum to s.
// Time complexity: O(n)
func Find(a []int, s int) (ps [][2]int) {
	ps = make([][2]int, 0)
	i, j := 0, len(a)-1

	for j > i {
		ai, aj := a[i], a[j]
		s1 := ai + aj
		if s1 == s {
			p := [2]int{ai, aj}
			ps = append(ps, p)
			ndup := 1
			for i++; j > i && a[i] == ai; i++ {
				ps = append(ps, p)
				ndup++
			}
			i = i - ndup
			j--
		} else if s1 > s {
			j--
		} else if s1 < s {
			i++
		}
	}

	return
}

// Finds all pairs of integers (duplicates included) in array that sum to s.
// Time complexitiy: O(n^2)
func FindBruteForce(a []int, s int) (ps [][2]int) {
	ps = make([][2]int, 0)

	for i := 0; i < len(a); i++ {
		ai := a[i]
		for j := i + 1; j < len(a); j++ {
			aj := a[j]
			s1 := ai + aj
			if s1 == s {
				p := [2]int{ai, aj}
				ps = append(ps, p)
			} else if s1 > s {
				break
			}
		}
	}

	return
}

// Check if two arrays of pairs are equal.
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

// Filter out all duplicate pairs.
// Assumes that pairs are sorted, i.e. there are
// no gaps between duplicates.
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
