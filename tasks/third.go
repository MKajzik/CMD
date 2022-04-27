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
	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "[ -d \"./path/dir/\" ] && echo \"Directory /path/dir/ exists.\"").Output()
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSuffix(string(out), "\n") == "Directory /path/dir/ exists." {
			text.Write("Brawo udało Ci się ukończyć drugie zadanie.")
			break
		}

		text.Write("Niestety popełniłeś gdzieś błąd. Spróbuj ponownie.")
	}

	time.Sleep(3 * time.Second)
	done <- true
}
