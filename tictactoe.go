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

	position := 1

	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLUMNS; j++ {
			board[ROWS*i+j] = '0' + rune(position)
			position++
		}
	}
}

func clearscreen() {

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
		for j := 0; j < COLUMNS; j++ {
			if j == 0 {
				fmt.Printf("\t")
			}
			fmt.Printf(" %c ", board[ROWS*i+j])
			if j < 2 {
				fmt.Printf("|")
			}
		}
		if i < 2 {
			fmt.Printf("\n\t-----------\n")
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
	fmt.Printf("Type a position: ")
	var position int

	_, err := fmt.Scan(&position)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for validation(position) == 0 {
		fmt.Printf("The position is not valid! Try again: ")
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
		fmt.Printf("Victory to the %c player!\n", player)
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
		fmt.Printf("Uh... It's a draw!")
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

func loop() {
	fmt.Println("Do you want to play again? ('1' to play again, '0' to leave)")
	var input int

	_, err3 := fmt.Scan(&input)
	if err3 != nil {
		fmt.Println("Error: ", err3)
		return
	}

	if input == 1 {
		main()
	} else {
		clearscreen()
		fmt.Println("Thanks for playing the game!")
	}

}

func main() {
	load()
	game()
	loop()
}
