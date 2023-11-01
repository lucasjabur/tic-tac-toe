package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const ROWS = 3
const COLUMNS = 3
const P1 = 'X'
const P2 = 'O'

var board [ROWS * COLUMNS]rune

func load() {
	// Inicializa o tabuleiro

	position := 0

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			board[ROWS*i+j] = rune(position)
			position++
		}
	}
}

func clearscreen() {
	// Irá limpar o terminal do jogador para inserção do tabuleiro

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func gameboard() {
	clearscreen()
	fmt.Printf("-=-=-=- Tic Tac Toe -=-=-=-\n\n")
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ { // inicialização do tabuleiro
			if j == 0 {
				fmt.Printf("\t") // tabulação do tabuleiro
			}
			fmt.Printf(" %c ", board[ROWS*i+j]) // print do tabuleiro
			if j < 2 {
				fmt.Printf(" |") // separação das colunas
			}
		}
		if i < 2 {
			fmt.Printf("\n\t-----------\n") // separação das linhas
		}
	}
	fmt.Printf("\n\n--=-=-=-=-=-=-=-=-=-=-=-=--")
	fmt.Printf("\n\n")
}

func validation(position int) int {
	empty := board[position-1] != P1 && board[position-1] != P2
	if position > 0 && position <= (ROWS*COLUMNS) && empty {
		return 1
	}
	return 0
}

func coordinates(player rune) {
	fmt.Println("Type a position: ")
	var position int

	_, err := fmt.Scan(&position)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for validation(position) == 0 {
		fmt.Println("The position is not valid! Try again: ")
		var position int

		_, err2 := fmt.Scan(&position)
		if err2 != nil {
			fmt.Println("Error: ", err2)
			return
		}
	}
	board[position-1] = player
}

func rows(player rune) bool {
	fstrow := board[0] == player && board[1] == player && board[2] == player
	sndrow := board[3] == player && board[4] == player && board[5] == player
	trdrow := board[6] == player && board[7] == player && board[8] == player

	if fstrow || sndrow || trdrow {
		return true
	}
	return false
}

func columns(player rune) bool {
	fstcolumn := board[0] == player && board[3] == player && board[6] == player
	sndcolumn := board[1] == player && board[4] == player && board[7] == player
	trdcolumn := board[2] == player && board[5] == player && board[8] == player

	if fstcolumn || sndcolumn || trdcolumn {
		return true
	}
	return false
}

func diagonal(player rune) bool {
	principal := board[0] == player && board[4] == player && board[8] == player
	secondary := board[2] == player && board[4] == player && board[6] == player

	if principal || secondary {
		return true
	}
	return false
}

func victory(player rune) bool {
	if rows(player) || columns(player) || diagonal(player) {
		fmt.Println("Victory to the", player, "!")
		return true
	}
	return false
}

func draw() bool {
	count := 0
	for i := 0; i < ROWS*COLUMNS; i++ {
		if board[i] == P1 || board[i] == P2 {
			count++
		}
	}
	if count == 9 {
		fmt.Println("Uh... It's a draw!")
		return true
	}
	return false
}

func game() {
	player := P1
	gameboard()
	for {
		coordinates(player)
		gameboard()
		if victory(player) || draw() {
			break
		}
		if player == P1 {
			player = P2
		} else {
			player = P1
		}
	}
}

func main() {
	load()
	game()
}
