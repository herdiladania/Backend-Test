package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	id    int
	dice  []int
	score int
}

func giveDiceToNextPlayer(players []Player, currentPlayer int) {
	if currentPlayer == len(players)-1 {
		players[0].dice = append(players[0].dice, 1)
	} else {
		players[currentPlayer+1].dice = append(players[currentPlayer+1].dice, 1)
	}
}

func removeDiceFromPlayer(players []Player, playerIndex, diceIndex int) {
	players[playerIndex].dice = append(players[playerIndex].dice[:diceIndex], players[playerIndex].dice[diceIndex+1:]...)
}

func addPointToPlayer(players []Player, playerIndex int) {
	players[playerIndex].score++
}

func getWinner(players []Player) int {
	maxScore := 0
	winner := -1
	for i, player := range players {
		if player.score > maxScore {
			maxScore = player.score
			winner = i
		}
	}
	return winner
}

func main() {
	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDice)

	players := make([]Player, numPlayers)
	for i := range players {
		players[i] = Player{
			id:    i + 1,
			dice:  make([]int, numDice),
			score: 0,
		}
	}

	rand.Seed(time.Now().UnixNano())

	round := 1
	lastPlayer := -1

	for {
		fmt.Println("==================")
		fmt.Printf("Giliran %d lempar dadu:\n", round)

		for i := range players {
			if len(players[i].dice) == 0 {
				continue
			}
			for j := range players[i].dice {
				players[i].dice[j] = rand.Intn(6) + 1
			}
			fmt.Printf("Pemain #%d (%d): %v\n", players[i].id, players[i].score, players[i].dice)
		}

		fmt.Printf("Setelah evaluasi:\n")
		for i := range players {
			if len(players[i].dice) == 0 {
				continue
			}
			for j := 0; j < len(players[i].dice); j++ {
				switch players[i].dice[j] {
				case 1:
					giveDiceToNextPlayer(players, i)
					removeDiceFromPlayer(players, i, j)
					j--
				case 6:
					addPointToPlayer(players, i)
					removeDiceFromPlayer(players, i, j)
					j--
				}
			}
			fmt.Printf("Pemain #%d (%d): ", players[i].id, players[i].score)
			if len(players[i].dice) == 0 {
				fmt.Println("_ (Berhenti bermain karena tidak memiliki dadu)")
			} else {
				fmt.Println(players[i].dice)
			}
		}

		activePlayers := 0
		lastPlayer = -1
		for i := range players {
			if len(players[i].dice) > 0 {
				activePlayers++
				lastPlayer = i
			}
		}
		if activePlayers == 1 {
			break
		}

		round++

	}

	fmt.Println("==================")
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", players[lastPlayer].id)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", getWinner(players)+1)
}
