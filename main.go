package main

import (
	"fmt"
	"github.com/astrieanna/tictactoe/tictactoe"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Server Starting")
	//http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		param := r.FormValue("board")
		board := tictactoe.FromString(param)

		// Check that the board was parse-able
		if board == nil {
			if len(param) != 9 {
				http.Error(w, "A tictactoe board should be 9 characters, but you sent me "+strconv.Itoa(len(param))+".", 400)
			} else {
				http.Error(w, "At least one character in your board was not 'ox '. You sent: "+param, 400)
			}
			return
		}

		// Check that it's our turn and the board isn't full
		err := board.Validate()
		if err != "" {
			http.Error(w, err, 400)
			return
		}

		// Check that the game isn't already over
		prewin := board.CheckForWin()
		if prewin != nil {
			if *prewin == tictactoe.EMPTY {
				// this would involve a full board, so shouldn't get here
				http.Error(w, "We're already in a draw", 400)
				log.Print("Saw a draw after checking for full board: ", param)
			} else if *prewin == tictactoe.O {
				http.Error(w, "I've already won", 400)
			} else if *prewin == tictactoe.X {
				http.Error(w, "You've already won", 400)
			}
		}

		board.MakeMove()
		fmt.Fprintf(w, board.ToString())

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
