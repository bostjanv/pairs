package pairs

import "testing"
import "time"
import "os"
import "math/rand"

const niter = 1000

func TestMain(m *testing.M) {
    rand.Seed(time.Now().UnixNano())
    os.Exit(m.Run())
}

func TestCornerCases(t *testing.T) {
    var a []int
    var ps [][2]int
    var ps1 [][2]int

    a = []int{1}
    ps = Find(a, 1)
    ps1 = [][2]int{}
    if CheckEqual(ps, ps1) == false {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    ps = FindUnique(a, 1)
    if CheckEqual(ps, ps1) == false {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    ps = FindBruteForce(a, 1)
    if CheckEqual(ps, ps1) == false {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }

    a = []int{1, 1}
    ps = Find(a, 2)
    ps1 = [][2]int{{1,1}}
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    ps = FindUnique(a, 2)
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    ps = FindBruteForce(a, 2)
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }

    a = []int{1, 1, 1}
    ps = Find(a, 2)
    ps1 = [][2]int{{1,1}, {1,1}, {1,1}}
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    ps = FindBruteForce(a, 2)
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    ps = FindUnique(a, 2)
    ps1 = [][2]int{{1,1}}
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
}

func TestRandom(t *testing.T) {
    for i := 0; i < 1000; i++ {
        n := rand.Intn(10000) + 1
        maxval := rand.Intn(10000)
        s := rand.Intn(maxval)
        runOnce(t, n, maxval, s)
    }
}

func runOnce(t *testing.T, n, maxval, s int) {
    a := MakeSortedArray(n, maxval)

    ps1 := Find(a, s)
    ps2 := FindBruteForce(a, s)

    if CheckEqual(ps1, ps1) != true {
        t.Errorf("ps1 != ps2")
    }

    ps3 := FindUnique(a, s)
    ps2 = FilterDup(ps2)

    if CheckEqual(ps2, ps3) != true {
        t.Errorf("ps2 != ps3")
    }

}

func BenchmarkFind(b *testing.B) {
    a := MakeSortedArray(10000, 10000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Find(a, 5000)
    }
}

func BenchmarkFindUnique(b *testing.B) {
    a := MakeSortedArray(10000, 10000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FindUnique(a, 5000)
    }
}

func BenchmarkFindBruteForce(b *testing.B) {
    a := MakeSortedArray(10000, 10000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FindBruteForce(a, 5000)
    }
}
