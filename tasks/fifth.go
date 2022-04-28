package tasks

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/bengadbois/flippytext"
)

func Fifth(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Piąte zadanie.\nW katalogu istnieje plik access.log. Policz linie w pliku które posiadają słowo \"GET\"")
	text.Write("Zapisz wynik w pliku result.txt")

	out, err := exec.Command("sh", "-c", "pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	text.Write(fmt.Sprintf("To zadanie wykonaj samodzielnie w oddzielnym terminalu zaczynając od komendy: cd %s", out))
	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "cat result.txt | tr -d \" \t\n\r\"").Output()
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSuffix(string(out), "\n") == "8" {
			text.Write("Brawo udało Ci się ukończyć trzecie zadanie.")
			err := os.Remove("result.txt")
			if err != nil {
				log.Fatal(err)
			}
			break
		}

		text.Write("Niestety popełniłeś gdzieś błąd. Spróbuj ponownie.")
	}

	time.Sleep(3 * time.Second)
	done <- true
}
