package behavior

import (
	"fmt"
	"math/rand"
	"time"
)

func TypeHuman(text string) {
	for _, ch := range text {
		fmt.Printf("%c", ch)

		delay := rand.Intn(200) + 50
		time.Sleep(time.Duration(delay) * time.Millisecond)

		if rand.Intn(10) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println()
}
