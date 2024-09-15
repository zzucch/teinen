package read

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zzucch/teinen/internal/waitlog"
)

func Read() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the input lines and a line 'end' to end:")

  lines := make([]string, 0)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
      return nil, err
		}

		if strings.TrimSpace(line) == "end" {
			break
		}

		lines = append(lines, line)
	}

	waitlog.Println("press enter to confirm")

	return lines, nil
}
