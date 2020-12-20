package heuristic

import (
	"math"
)

type LateAcceptance struct {
	n       int
	tScore  []float64
}

func (l *LateAcceptance) Init(n int) {
	l.n = n + 1
	l.tScore = append(l.tScore, math.MaxFloat64)
}

func (l *LateAcceptance) Accept(pScore float64) bool {

	var pAccept bool
	if pScore < l.tScore[0] {
		pAccept = true
	} else if pScore < l.tScore[len(l.tScore) - 1] {
		pAccept = true
	} else {
		pAccept = false
	}
	l.tScore = append(l.tScore, pScore)
	//fmt.Println(l.tScore)
	if len(l.tScore) > l.n {
		l.tScore = l.tScore[1:]
	}
	return pAccept
}
