package tasks

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/bengadbois/flippytext"
)

func Third(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Trzecie zadanie.\nUtwórz katalog \"tmp/files/secret\". \nPamiętaj, że foldery tmp i files nie istnieją i musisz je osobno stworzyć.")
	text.Write("Podpowiedź: Da się to zrobić jedną komendą")

	out, err := exec.Command("sh", "-c", "pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	text.Write(fmt.Sprintf("To zadanie wykonaj samodzielnie w oddzielnym terminalu zaczynając od komendy: cd %s", out))

	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "[ -d \"./tmp/files/secret/\" ] && echo \"Directory /tmp/files/secret/ exists.\"").Output()
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSuffix(string(out), "\n") == "Directory /tmp/files/secret/ exists." {
			text.Write("Brawo udało Ci się ukończyć trzecie zadanie.")
			break
		}

		text.Write("Niestety popełniłeś gdzieś błąd. Spróbuj ponownie.")
	}

	time.Sleep(3 * time.Second)
	done <- true
}
