package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Print("$ ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		out, err := exec.Command("sh", "-c", input).Output()
		if err != nil {
			log.Fatal(err)
		}
		text.Write(fmt.Sprintf("Output: %s", out))

		if strings.TrimSuffix(input, "\n") == "pwd" {
			text.Write("Brawo udało Ci się ukończyć drugie zadanie.")
			break
		}

		text.Write("Niestety popełniłeś gdzieś błąd. Spróbuj ponownie.")
	}

	time.Sleep(3 * time.Second)
	done <- true
}
