package runner

import (
	"bufio"
	"fmt"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/pwnesia/dnstake/internal/errors"
)

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

func read(file *os.File) ([]string, error) {
	var lines []string

	keys := make(map[string]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		host := scanner.Text()
		if _, value := keys[host]; !value {
			if govalidator.IsHost(host) {
				keys[host] = true
				lines = append(lines, host)
			}
		}
	}

	if len(lines) < 1 {
		return lines, fmt.Errorf("%s", errors.ErrNoValid)
	}

	return lines, scanner.Err()
}
