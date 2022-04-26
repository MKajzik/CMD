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

func First(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Pierwsze zadanie.\nWyświetl tekst \"hello world\" używając tylko jednej komendy.")
	fmt.Print("$ ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	out, err := exec.Command("sh", "-c", input).Output()
	if err != nil {
		log.Fatal(err)
	}
	text.Write(fmt.Sprintf("Output: %s", out))

	if strings.TrimSuffix(string(out), "\n") == "hello world" {
		text.Write("Brawo udało Ci się ukończyć pierwsze zadanie.")
	}

	time.Sleep(3 * time.Second)
	done <- true
}
