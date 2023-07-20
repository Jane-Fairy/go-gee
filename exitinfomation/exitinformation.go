package exitinfomation

import "math/rand"

//随机数生成，等待通知结果
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for true {
			select {
			case ch <- rand.Int():
			case <-done:
				break Label
			}
		}

		close(ch)
	}()

	return ch
}
