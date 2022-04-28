package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"warsztaty/aplikacja/tasks"

	"github.com/bengadbois/flippytext"
)

func main() {
	clearTerminal()
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Witaj w naszym wyzwaniu CMD! \nZa chwilę dostaniesz 8 zadań do wykonania aby zaliczyć nasze wyzwanie.")
	text.Write("Naciśnij ENTER kiedy będziesz gotowy")
	fmt.Scanln()
	done := make(chan bool, 1)
	clearTerminal()
	go tasks.First(done)
	<-done
	clearTerminal()
	go tasks.Second(done)
	<-done
	clearTerminal()
	text.Write("Rozgrzewka skończona czas przejść do bardziej ciekawych wyzwań!")
	go tasks.Third(done)
	<-done
	clearTerminal()
	go tasks.Fourth(done)
	<-done
	clearTerminal()
	go tasks.Fifth(done)
	<-done
	clearTerminal()
	go tasks.Sixth(done)
	<-done
	clearTerminal()
	go tasks.Seventh(done)
	<-done
	clearTerminal()
	go tasks.Eigth(done)
	<-done
	text.Write("GRATULACJE! Udało Ci się ukończyć nasze wyzwanie CMD.")
	time.Sleep(5 * time.Second)
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
