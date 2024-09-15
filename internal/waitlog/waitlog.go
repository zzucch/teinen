package waitlog

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Println(s string) {
	waitForInputAfter(func() {
		fmt.Println(s)
	})
}

func Fatal(a ...any) {
	waitForInputAfter(func() {
		fmt.Println(a...)
	})

	os.Exit(1)
}

func waitForInputAfter(fn func()) {
	reader := bufio.NewReader(os.Stdin)

	fn()

	_, _, err := reader.ReadLine()
	if err != nil {
    log.Fatal(err)
	}
}
