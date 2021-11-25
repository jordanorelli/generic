package iter

import (
	"testing"
	"strconv"
)

func TestSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	s := Slice(nums)
	for n, it := Start(s); it.Next(&n); {
		t.Log(n)
	}


	for n, it := Start(Map(Slice(nums), strconv.Itoa)); it.Next(&n); {
		t.Log("fart " + n)
	}
}
