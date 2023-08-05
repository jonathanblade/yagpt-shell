package chat

import (
	"fmt"
	"strings"
	"time"
)

func typing(s string) {
	parts := strings.Split(s, "\n")
	for _, part := range parts {
		for c := 0; c < len(part); c++ {
			fmt.Printf("\r%s", part[:c])
			time.Sleep(25 * time.Millisecond)
		}
		fmt.Print("\n")
	}
}
