package exitinfomation

import "testing"

func TestChina(t *testing.T) {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	china := China(China(in))
	for v := range china {
		println(v)
	}
}
