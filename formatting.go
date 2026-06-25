package main

import "fmt"

func main() {
	playerName := "Ada"
	level := 7
	guild := "Rangers"
	isPremium := true

	fmt.Printf("Player: %s\n", playerName)
	fmt.Printf("Level: %d\n", level)
	fmt.Printf("Guild: %s\n", guild)
	fmt.Printf("Premium: %t\n", isPremium)
}

/*output:
Player: Ada
Level: 7
Guild: Rangers
Premium: true*/
