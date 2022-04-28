package tasks

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"warsztaty/aplikacja/shared"

	"github.com/bengadbois/flippytext"
)

func Nine(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Dziewiąte zadanie.\nW katalogu \"nine/\" zostały dodane pliki .txt. Za pomocą komendy usuń z plików słowo \"challenge\" rekursywnie")
	createDir9()

	out, err := exec.Command("sh", "-c", "pwd").Output()
	if err != nil {
		log.Fatal(err)
	}

	text.Write(fmt.Sprintf("To zadanie wykonaj samodzielnie w oddzielnym terminalu zaczynając od komendy: cd %s", out))

	for {
		text.Write("Naciśnij ENTER kiedy wykonasz zadanie")
		fmt.Scanln()

		out, err := exec.Command("sh", "-c", "./script.sh").Output()
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSuffix(string(out), "\n") == "NOT OK" {
			text.Write("Brawo udało Ci się ukończyć dziewiąte zadanie.")
			err := os.RemoveAll("nine/")
			if err != nil {
				log.Fatal(err)
			}
			err = os.Remove("script.sh")
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

func createDir9() {
	if err := os.RemoveAll("nine/"); err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir("nine", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	err := os.WriteFile("script.sh", []byte(shared.Script), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("nine/aut.txt", []byte(shared.Lorem), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("nine/deleniti.txt", []byte("challenge "+shared.Lorem), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("nine/dolorem.txt", []byte(shared.Lorem+"challenge"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("nine/dolorum.txt", []byte(shared.Lorem+"Linux is great"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("nine/hic.txt", []byte(shared.Lorem), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = os.WriteFile("nine/id.txt", []byte(shared.Lorem+"Linux is challenge"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
