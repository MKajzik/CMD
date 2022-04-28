package tasks

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"warsztaty/aplikacja/shared"

	"github.com/bengadbois/flippytext"
)

func Sixth(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Piąte zadanie.\nW katalogu istnieje plik users.log. Usuń z każdej nazwy użytkownika znak \"_\"")
	text.Write("Zapisz wynik w pliku result.txt")

	out, err := exec.Command("sh", "-c", "pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	text.Write(fmt.Sprintf("To zadanie wykonaj samodzielnie w oddzielnym terminalu zaczynając od komendy: cd %s", out))

	err = os.WriteFile("users.log", []byte(shared.Six), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "cat result.txt").Output()
		if err != nil {
			log.Fatal(err)
		}

		if string(out) == shared.Six_solution {
			text.Write("Brawo udało Ci się ukończyć szóste zadanie.")
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
