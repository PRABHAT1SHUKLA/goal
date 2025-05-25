package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	Empty = 0
	AI    = 1
	Human = 2
	Size  = 9
	Win   = 5
)

type Move struct {
	Row int
	Col int
}

var board [Size][Size]int

func main() {
	rand.Seed(time.Now().UnixNano())
	board[4][4] = Human // human move
	move := bestMove(AI, 2)
	fmt.Printf("AI plays at: (%d, %d)\n", move.Row, move.Col)
	board[move.Row][move.Col] = AI
	printBoard()
}

func bestMove(player int, depth int) Move {
	best := Move{-1, -1}
	bestScore := math.Inf(-1)

	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if board[r][c] == Empty {
				board[r][c] = player
				score := minimax(depth-1, false, player)
				board[r][c] = Empty

				if score > bestScore {
					bestScore = score
					best = Move{r, c}
				}
			}
		}
	}
	return best
}

func minimax(depth int, maximizing bool, player int) float64 {
	winner := checkWinner()
	if winner == player {
		return 1000
	} else if winner != 0 {
		return -1000
	}
	if depth == 0 {
		return float64(evaluate(player))
	}

	bestScore := math.Inf(-1)
	if !maximizing {
		bestScore = math.Inf(1)
	}

	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if board[r][c] == Empty {
				board[r][c] = player
				score := minimax(depth-1, !maximizing, 3-player)
				board[r][c] = Empty

				if maximizing {
					bestScore = math.Max(bestScore, score)
				} else {
					bestScore = math.Min(bestScore, score)
				}
			}
		}
	}
	return bestScore
}

func evaluate(player int) int {
	score := 0
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if board[r][c] == player {
				score += countDirection(r, c, 1, 0, player)  // horizontal
				score
