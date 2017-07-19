package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				// this doesn't work
				// f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				// this does
				f, _ := os.OpenFile("./log.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339)
				n, err := f.WriteString(logTime + " - " + msg)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("wrote ", n, " bytes")
				}
				f.Close()
			} else {
				break
			}
		}
	}()

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			go func(i, j int) {
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}

	fmt.Scanln()
}
