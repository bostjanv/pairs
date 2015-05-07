package pairs

import "testing"
import "time"
import "os"
import "math/rand"

func TestMain(m *testing.M) {
    rand.Seed(time.Now().UnixNano())
    os.Exit(m.Run())
}

func TestCornerCases(t *testing.T) {
    var a []int
    var ps [][2]int
    var ps1 [][2]int
    var c uint

    a = []int{1}
    ps, c = Find(a, 1)
    ps1 = [][2]int{}
    if CheckEqual(ps, ps1) == false {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    if c != 0 {
        t.Errorf("c == %d, want %d", c, 0)
    }

    a = []int{1, 1}
    ps, c = Find(a, 2)
    ps1 = [][2]int{{1,1}}
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    if c != 1 {
        t.Errorf("c == %d, want %d", c, 1)
    }

    a = []int{1, 1, 1}
    ps, c = Find(a, 2)
    ps1 = [][2]int{{1,1}, {1,1}, {1,1}}
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    if c != 3 {
        t.Errorf("c == %d, want %d", c, 3)
    }

    ps, c = FindUnique(a, 2)
    ps1 = [][2]int{{1,1}}
    if CheckEqual(ps, ps1) == false  {
        t.Errorf("ps == %d, want %d", ps, ps1)
    }
    //if c != 1 {
    //    t.Errorf("c == %d, want %d", c, 1)
    //}

}

func TestRandom(t *testing.T) {
    foo(t, 10000, 10000, 100, 100)
    foo(t, 10, 10000, 100, 100)
    foo(t, 1000, 1, 100, 100)
    foo(t, 1000, 1, 100, 100)
}

func foo(t *testing.T, n, max, iter, s int) {
    var a []int
    var ps1, ps2, ps3 [][2]int

    for i := 0; i < iter; i++ {
        a = MakeSortedArray(n, max)

        ps1, _ = Find(a, s)
        ps2, _ = FindBruteForce(a, 100)

        if CheckEqual(ps1, ps1) != true {
            t.Errorf("ps1 != ps2")
        }

        ps3, _ = FindUnique(a, 100)
        ps2 = FilterDup(ps2)

        if CheckEqual(ps2, ps3) != true {
            t.Errorf("ps2 != ps3")
        }
    }

}
