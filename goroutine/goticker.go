package goroutine

import (
	"fmt"
	"time"
)

func gotickerTest() {
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				//停止
				ticker.Stop()
			}
		}
	}()
	for {
	}
}