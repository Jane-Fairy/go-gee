package contextTest

import (
	"context"
	"fmt"
	"time"
)

func init() {
	con()
}

type otherContext struct {
	context.Context
}

func con() {
	ctxa, cancel := context.WithCancel(context.Background())
	go work(ctxa, "work1")

	tm := time.Now().Add(2 * time.Second)
	ctxb, _ := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work2")

	oc := otherContext{ctxb}

	ctxc := context.WithValue(oc, "key", "aandes , pass from main.")

	go workWithValue(ctxc, "work3")

	time.Sleep(10 * time.Second)

	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("main.stop")

}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return

		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}

	}

}

func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel \n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value = %s \n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}
