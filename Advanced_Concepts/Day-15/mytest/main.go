package main

import (
	"mytest/player"
	"fmt"
)

func main() {
	s := player.Stats{Name: "Player1", Minutes: 25.1, Points: 21, Assist: 3, TurnOver: 7, Rebounds: 3}
	fmt.Println(player.HadAGoodGame(s))
}