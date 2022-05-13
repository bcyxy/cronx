package gfunc_test

import (
	"fmt"
	"testing"

	"github.com/bcyxy/cronx/common/gfunc"
)

func TestXxx(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(gfunc.GetUniqID())
	}
}
