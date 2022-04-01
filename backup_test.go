// package backup

// import (
// 	"testing"
// )

// func TestSaveAndLoad(t *testing.T) {
// 	cases := []struct {
// 		in, want []bool
// 	}{
// 		{[]bool{true, true, true, true}, []bool{true, true, true, true}},
// 		{[]bool{false, true, false, true}, []bool{false, true, false, true}},
// 		{[]bool{false, false, true, true}, []bool{false, false, true, true}},
// 		{[]bool{false, false, false, false}, []bool{false, false, false, false}},
// 	}
// 	for _, c := range cases {
// 		SaveCab(c.in)
// 		got := LoadCab("orders.txt")
// 		for i := 0; i < len(got); i++ {
// 			if got[i] != c.want[i] {
// 				t.Errorf("Error:\n In:   %v --> Out: %v\n Want: %v", c.in, got, c.want)
// 			}
// 		}
// 	}
// }
