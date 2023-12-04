package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	timer := time.NewTimer(10 * time.Second)

	startTime := time.Now()

	go func() {
		<-timer.C
		fmt.Println("Goodbay World!")
		os.Exit(0)
	}()
	fmt.Scanln()

	stop := timer.Stop()
	if stop {
		fmt.Println("Stopped by the user after", time.Since(startTime), "seconds!")
	}

}
