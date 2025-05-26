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
				score += countDirection(r, c, 0, 1, player)  // vertical
				score += countDirection(r, c, 1, 1, player)  // diagonal /
				score += countDirection(r, c, 1, -1, player) // diagonal \
			}
		}
	}
	return score
}

func countDirection(r, c, dr, dc, player int) int {
	count := 0
	for i := 0; i < Win && r+i*dr < Size && c+i*dc < Size && r+i*dr >= 0 && c+i*dc >= 0; i++ {
		if board[r+i*dr][c+i*dc] == player {
			count++
		} else {
			break
		}
	}
	return count * count
}

func checkWinner() int {
	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			if board[r][c] == Empty {
				continue
			}
			player := board[r][c]
			if checkDir(r, c, 1, 0, player) || checkDir(r, c, 0, 1, player) || checkDir(r, c, 1, 1, player) || checkDir(r, c, 1, -1, player) {
				return player
			}
		}
	}
	return 0
}

func checkDir(r, c, dr, dc, player int) bool {
	for i := 0; i < Win; i++ {
		nr := r + i*dr
		nc := c + i*dc
		if nr < 0 || nr >= Size || nc < 0 || nc >= Size || board[nr][nc] != player {
			return false
		}
	}
	return true
}

func printBoard() {
	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case Empty:
				fmt.Print(". ")
			case AI:
				fmt.Print("O ")
			case Human:
				fmt.Print("X ")
			}
		}
		fmt.Println()
	}
}
