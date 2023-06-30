package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	dir := "audio"
	err := os.RemoveAll(dir)
	if err != nil {
		fmt.Println(err)
	}
	speech := htgotts.Speech{
		Folder:   "audio",
		Language: voices.English,
	}
	speech.Speak("Ho is Navneet Shukla")

	latestFile := getLatestMP3File("audio") // Retrieve the latest generated audio file

	f, err := os.Open(latestFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		panic(err)
	}

	c, err := oto.NewContext(int(d.SampleRate()), 2, 2, 8192)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	fmt.Printf("Length: %d[bytes]\n", d.Length())

	if _, err := io.Copy(p, d); err != nil {
		log.Fatal(err)
	}
}

func getLatestMP3File(dirPath string) string {
	var latestFile string
	var latestModTime int64

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".mp3" && info.ModTime().UnixNano() > latestModTime {
			latestFile = path
			latestModTime = info.ModTime().UnixNano()
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return latestFile
}
