package math

import (
	"testing"
)

var combinationTests = []struct {
	n   int
	k   int
	ans int
	pan bool
}{
	{n: 5, k: 2, ans: 10},
	{n: 5, k: 0, ans: 1},
	{n: 5, k: 3, ans: 10},
	{n: 5, k: 5, ans: 1},
	{n: 0, k: 0, ans: 1},
	{n: 10, k: 0, ans: 1},
	{n: 10, k: 1, ans: 10},
	{n: 10, k: 2, ans: 45},
	{n: 10, k: 3, ans: 120},
	{n: 10, k: 4, ans: 210},
	{n: 10, k: 5, ans: 252},
	{n: 10, k: 6, ans: 210},
	{n: 10, k: 7, ans: 120},
	{n: 10, k: 8, ans: 45},
	{n: 10, k: 9, ans: 10},
	{n: 10, k: 10, ans: 1},
	{n: 1000, k: 500, pan: true},
}

func TestMath(t *testing.T) {
	for _, ts := range combinationTests {
		if ts.pan {
			if !combinationPanics(ts.n, ts.k) {
				t.Errorf("Combination does not panic for n=%v, k=%v", ts.n, ts.k)
			}
			continue
		}
		ans := Combination(ts.n, ts.k)
		if ans != ts.ans {
			t.Errorf("Combination, n=%v, k=%v: got %v, expected %v", ts.n, ts.k, ans, ts.ans)
		}
	}
}

func combinationPanics(n, k int) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			b = true
		}
	}()
	Combination(n, k)
	return
}
