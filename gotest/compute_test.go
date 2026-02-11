package main 

import "testing"

type test struct {
	nb,cubenb int 
}
func TestCube(t *testing.T) {

	var cube1 test 
	cube1.nb =5 
	cube1.cubenb = 125

    var cube2 test 
    cube2.nb =1
    cube2.cubenb =1 

	tests := []test{cube1,cube2}

	for _,tt := range tests {
		result := Cube(tt.nb)

		if result != tt.cubenb {
			t.Errorf(" INPUT %d  WRONG OUTPUT %d ",tt.nb,tt.cubenb)

		} else {
			t.Logf("CORRECT INPUT %d OUTPUT %d ",tt.nb,tt.cubenb)
		}



	}

    // call compute() and check results
}