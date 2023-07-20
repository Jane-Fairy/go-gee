package fanin

import "math/rand"

// GenerateIntA done for receive exit sign
func GenerateIntA(done chan struct{}) chan int {
	mych := make(chan int, 5)
	go func() {
	Label:
		for {
			select {
			case mych <- rand.Int():
			case <-done:
				break Label
			}
		}
		close(mych)
	}()
	return mych
}

//GenerateIntB done for receive exit sign
func GenerateIntB(done chan struct{}) chan int {

	mych := make(chan int, 10)

	go func() {
	Lable:
		for true {
			select {
			case mych <- rand.Int():
			case <-done:
				break Lable
			}
		}

		close(mych)
	}()
	return mych
}
