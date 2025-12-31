package behavior

import (
	"fmt"
	"math/rand"
	"time"
)

func MoveMouseHuman(targetX, targetY int) {
	x, y := 0, 0
	steps := rand.Intn(15) + 10

	for i := 0; i < steps; i++ {
		x += (targetX - x) / (steps - i)
		y += (targetY - y) / (steps - i)

		fmt.Printf("Mouse at (%d, %d)\n", x, y)

		time.Sleep(time.Duration(rand.Intn(40)+20) * time.Millisecond)
	}


	time.Sleep(200 * time.Millisecond)
}
