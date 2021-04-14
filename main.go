package main

import (
	"domino/piece"
	"domino/player"
	"fmt"
)

func main() {
	var numberOfPlayers int
	numberOfPlayers = 3
	players := player.GeneratePlayers(numberOfPlayers)
	for index,player := range players {
		fmt.Printf("\nPLayer %v\n",index)
		for _,piece := range player.Pieces {
			fmt.Printf("First half :- %v , Second Half :- %v\n",piece.FirstHalf,piece.SecondHalf)
		}
		fmt.Printf("Max Double :- %v\n",player.MaxDouble)
	}
	remainingPieces := piece.GenerateRemainingPieces(numberOfPlayers)
	fmt.Println("Player with bigger double :-  ",players.GetPlayerWithBiggerDouble().MaxDouble)
	fmt.Println(remainingPieces)
}
