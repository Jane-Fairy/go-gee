package exitinfomation

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGenerateIntA(t *testing.T) {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println("ssss", <-ch)
	fmt.Println(<-ch)

	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	println("NumberGoroutine=", runtime.NumGoroutine())

}
