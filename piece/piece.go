package piece

import (
	"domino/constants"
	"math/rand"
	"time"
)

type Piece struct {
	FirstHalf  int
	SecondHalf int
}

func GenerateSamplePiece() *Piece {
	rand.Seed(time.Now().UnixNano())
	return &Piece{
		FirstHalf:  rand.Intn(constants.MAX_DOT_VALUE-constants.MIN_DOT_VALUE) + constants.MIN_DOT_VALUE,
		SecondHalf: rand.Intn(constants.MAX_DOT_VALUE-constants.MIN_DOT_VALUE) + constants.MIN_DOT_VALUE,
	}
}

func GenerateRemainingPieces(numberOfPlayers int) []*Piece {
	remainingPieces := 28 - (7 * numberOfPlayers)
	var pieces []*Piece
	if remainingPieces > 0 {
		for i:=1 ; i<=remainingPieces ; i++ {
			pieces = append(pieces,GenerateSamplePiece())
		}
	}
	return pieces
}