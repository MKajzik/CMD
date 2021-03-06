package tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/bengadbois/flippytext"
)

func Fourth(done chan bool) {
	text := flippytext.New()
	text.TickerTime = time.Millisecond * 2
	text.Write("Czwarte zadanie.\nW folderze gdzie uruchomiono tą aplikację został stworzony plik \"access.log\"")
	text.Write("Wyświetl wszystkie linijki które posiadają w sobie \"GET\"")
	err := os.WriteFile("access.log", []byte(`163.56.115.58 - - [09/Jan/2017:22:29:57 +0100] "GET /posts/2/display HTTP/1.0" 200 3240
75.113.188.234 - - [09/Jan/2017:22:30:43 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 200 1116
69.16.40.148 - - [09/Jan/2017:22:34:33 +0100] "GET /pages/create HTTP/1.0" 500 3471
225.219.54.140 - - [09/Jan/2017:22:35:30 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 500 2477
207.243.19.2 - - [09/Jan/2017:22:38:03 +0100] "GET /bar/create HTTP/1.0" 200 1116
199.37.62.156 - - [09/Jan/2017:22:42:18 +0100] "GET /posts/1/display HTTP/1.0" 200 2477
55.74.240.123 - - [09/Jan/2017:22:44:25 +0100] "POST /posts/1/display HTTP/1.0" 200 3471
251.111.109.143 - - [09/Jan/2017:22:49:02 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 200 2477
101.163.230.250 - - [09/Jan/2017:22:52:31 +0100] "DELETE /posts/2/display HTTP/1.0" 404 2477
200.19.168.148 - - [09/Jan/2017:22:57:11 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 200 3471`), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	for {
		fmt.Print("$ ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		out, err := exec.Command("sh", "-c", input).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Output: %s", out)

		if string(out) == `163.56.115.58 - - [09/Jan/2017:22:29:57 +0100] "GET /posts/2/display HTTP/1.0" 200 3240
75.113.188.234 - - [09/Jan/2017:22:30:43 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 200 1116
69.16.40.148 - - [09/Jan/2017:22:34:33 +0100] "GET /pages/create HTTP/1.0" 500 3471
225.219.54.140 - - [09/Jan/2017:22:35:30 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 500 2477
207.243.19.2 - - [09/Jan/2017:22:38:03 +0100] "GET /bar/create HTTP/1.0" 200 1116
199.37.62.156 - - [09/Jan/2017:22:42:18 +0100] "GET /posts/1/display HTTP/1.0" 200 2477
251.111.109.143 - - [09/Jan/2017:22:49:02 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 200 2477
200.19.168.148 - - [09/Jan/2017:22:57:11 +0100] "GET /posts/foo?appID=xxxx HTTP/1.0" 200 3471
` {
			text.Write("Brawo udało Ci się ukończyć czwarte zadanie.")
			break
		}

		text.Write("Niestety popełniłeś gdzieś błąd. Spróbuj ponownie.")
	}

	time.Sleep(3 * time.Second)
	done <- true
}
