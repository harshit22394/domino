package player

import (
	"domino/piece"
	"sort"
)

type Players []*Player

type Player struct {
	Pieces []*piece.Piece
	MaxDouble int
}

func GeneratePlayers(n int) Players {
	var ps Players

	for i:=1 ; i<=n ; i++ {
		var p Player
		for j:=0 ; j<7 ;j++ {
			domino := piece.GenerateSamplePiece()
			p.Pieces = append(p.Pieces,domino)
			double := domino.FirstHalf + domino.SecondHalf
			if double > p.MaxDouble {
				p.MaxDouble = double
			}
		}
		ps = append(ps,&p)
	}
	return ps
}

func (ps Players) GetPlayerWithBiggerDouble() *Player {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].MaxDouble > ps[j].MaxDouble
	})
	return ps[0]
}