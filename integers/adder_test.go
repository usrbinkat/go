package integers

import (
	"fmt"
    "testing"
)

func TestAdder(t *testing.T) {

	expect := 4
	sum := Add(2, 2)

	if sum != expect {
		msg := "\nExpected SUM: %d \nReturned SUM: %d"
		t.Errorf(msg, expect, sum)
	}
	fmt.Println(sum)
}