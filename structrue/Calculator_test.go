package structrue

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	fmt.Println(calculate("1+(2*(2+1)+3/(2+8))"))
}
