package behavior

import (
	"fmt"
	"math/rand"
	"time"
)

func ScrollHuman() {
	scrollPosition := 0

	for i := 0; i < 6; i++ {

		move := rand.Intn(300) + 100
		scrollPosition += move

		fmt.Println(" Scrolled to:", scrollPosition)

		
		time.Sleep(time.Duration(rand.Intn(1200)+500) * time.Millisecond)

	
		if rand.Intn(4) == 0 {
			scrollPosition -= rand.Intn(100)
			fmt.Println(" Scrolled back to:", scrollPosition)
			time.Sleep(400 * time.Millisecond)
		}
	}
}
