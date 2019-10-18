package progress

import (
	"fmt"
	"time"
)

func SixDotLoop(end <-chan struct{}, startInfo string) {

	isDone := false
	go func() {
		<-end
		isDone = true
	}()
	for {
		fmt.Printf("\r%s", startInfo)
		time.Sleep(time.Second)
		for i := 0; i <= 5 ; i++  {
			if isDone {
				fmt.Printf(" ")
				return
			}
			fmt.Print(".")
			time.Sleep(time.Second)
		}
	}
}
