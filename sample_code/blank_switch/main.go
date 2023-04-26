package main

import (
	"fmt"
	"math/rand"
)

func main() {
	basicBlankSwitch()
	fmt.Println()
	overdoneBlankSwitch()
}

func basicBlankSwitch() {
	fmt.Println("a basic blank switch")
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}
}

func overdoneBlankSwitch() {
	fmt.Println("this blank switch would be better implemented as a regular switch")
	for i := 0; i < 10; i++ {
		a := rand.Intn(10)
		switch {
		case a == 2:
			fmt.Println("a is 2")
		case a == 3:
			fmt.Println("a is 3")
		case a == 4:
			fmt.Println("a is 4")
		default:
			fmt.Println("a is ", a)
		}
	}

	fmt.Println("\njust use a regular switch like this")
	for i := 0; i < 10; i++ {
		a := rand.Intn(10)
		switch a {
		case 2:
			fmt.Println("a is 2")
		case 3:
			fmt.Println("a is 3")
		case 4:
			fmt.Println("a is 4")
		default:
			fmt.Println("a is ", a)
		}
	}

}
