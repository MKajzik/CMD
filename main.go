package main

import (
	//"time"

	"time"

	"github.com/bengadbois/flippytext"
)

func main() {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 4
	text.Write("Witaj w naszym wyzwaniu CMD! \nZa chwilę dostaniesz 10 zadań do wykonania aby zaliczyć nasze wyzwanie.")
}
