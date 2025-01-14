package recover

import (
	"fmt"
	"testing"
)

func test() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0

	num1 = num1 / num2
	fmt.Println(num1)
}

func TestRecover(t *testing.T) {
	test()
	t.Error("下面的代码")
}
