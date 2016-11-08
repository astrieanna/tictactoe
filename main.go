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
		if board == nil {
			if len(param) != 9 {
				http.Error(w, "A tictactoe board should be 9 characters, but you sent me "+strconv.Itoa(len(param))+".", 400)
			} else {
				http.Error(w, "At least one character in your board was not 'ox '. You sent: "+param, 400)
			}
			return
		}
		err := board.Validate()
		if err != "" {
			http.Error(w, err, 400)
			return
		}
		board.MakeMove()
		fmt.Fprintf(w, board.ToString())

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
