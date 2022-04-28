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

func Seventh(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Siódme zadanie.\nW katalogu \"siedem/\" zostały dodane pliki. Usuń wszystkie pliki oraz wszystkie podkatalogi.")
	text.Write("Uwaga: istnieją pliki ukryte z \".\" . rm -rf * tutaj nie zadziała")
	createDir()

	out, err := exec.Command("sh", "-c", "pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	text.Write(fmt.Sprintf("To zadanie wykonaj samodzielnie w oddzielnym terminalu zaczynając od komendy: cd %s", out))

	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "find siedem -maxdepth 0 -empty -exec echo {} is empty. \\;").Output()
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSuffix(string(out), "\n") == "siedem is empty." {
			text.Write("Brawo udało Ci się ukończyć siódme zadanie.")
			err := os.RemoveAll("siedem/")
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

func createDir() {
	if err := os.RemoveAll("siedem/"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("siedem", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	err := os.WriteFile("siedem/.dont.forget.dotfiles", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("siedem/aut.webm", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("siedem/deleniti.gif", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("siedem/dolorem.doc", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("siedem/dolorum.wav", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("siedem/hic.xls", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("siedem/id.txt", []byte("1"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
