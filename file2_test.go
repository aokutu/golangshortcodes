package main

import "testing"

type testtools struct {
	a, b, c float64
}

func Triarea(a, b float64) float64 {
	return 0.5 * a * b
}

func TestTriarea(t *testing.T) {
	var test1 testtools
	test1.a, test1.b, test1.c = 5, 4, 10
	var test2 testtools
	test2.a, test2.b, test2.c = 6, 6, 18
	var test3 testtools
	test3.a, test3.b, test3.c = 10, 10, 50

	tests := []testtools{test1, test2, test3}

	for _, tt := range tests {
		tst := Triarea(tt.a, tt.b)
		if tst != tt.c {
			//t.Errorf("Triarea(%f, %f) = %f; expected %f", tt.a, tt.b, tst, tt.c)

			t.Errorf("WRONG OUTCOME EXPECTED %f BUT GOT %f",tt.c,tst)
		} else {
		t.Logf("SUCCESS: TRIAREA(%f, %f) EXPECTED %f OUTPUT %f", tt.a, tt.b ,tt.c,tst)
	} 
	}
}
