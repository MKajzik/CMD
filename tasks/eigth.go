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

func Eigth(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Ósme zadanie.\nW katalogu \"osiem/\" zostały dodane pliki. Za pomocą komendy zapisz wielkość plików w katalogu osiem do pliku result.txt")
	createDir8()

	out, err := exec.Command("sh", "-c", "pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	text.Write(fmt.Sprintf("To zadanie wykonaj samodzielnie w oddzielnym terminalu zaczynając od komendy: cd %s", out))

	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "cat result.txt").Output()
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSuffix(string(out), "\n") == `
6
16
4
21
26
52` {
			text.Write("Brawo udało Ci się ukończyć ósme zadanie.")
			err := os.RemoveAll("osiem/")
			if err != nil {
				log.Fatal(err)
			}
			err = os.Remove("result.txt")
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

func createDir8() {
	if err := os.RemoveAll("osiem/"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("osiem", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	err := os.WriteFile("osiem/aut.webm", []byte("111111"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("osiem/deleniti.gif", []byte("1111112123123123"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("osiem/dolorem.doc", []byte("1111"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("osiem/dolorum.wav", []byte("112413451345123541235"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("osiem/hic.xls", []byte("11324512351235123512351235"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("osiem/id.txt", []byte("1123514325135146134gdfghsdfbadfhaergasdfgdsrhasdgasd"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
