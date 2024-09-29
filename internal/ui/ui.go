package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zzucch/teinen/internal/waitlog"
)

func Choose(choosing string, options *[]string) string {
	fmt.Println(choosing + " options:")
	for i, option := range *options {
		fmt.Println(strconv.Itoa(i+1) + ": " + option)
	}

	fmt.Println("enter the chosen number")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		waitlog.Fatal("Error reading input")
	}

	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(*options) {
		waitlog.Fatal("Invalid choice")
	}

	return (*options)[choice-1]
}
