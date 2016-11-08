package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/astrieanna/tictactoe/tictactoe"
)

func main() {
	fmt.Println("hello world")
	//http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		param := r.FormValue("board")
		board := tictactoe.FromString(param)
		fmt.Fprintf(w, board.ToString())

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
