package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lofi-cli <play>")
		return
	}

	if os.Args[1] == "play" {
		playInBrowser()
	}
}

func playInBrowser() {
	// These can be YouTube links or any streaming URLs
	tracks := []string{
		"https://www.youtube.com/watch?v=AMcVJmb5mvk&t=4086s&pp=ygUKdG9reW8gbG9maQ%3D%3D",
	}

	rand.Seed(time.Now().UnixNano())
	url := tracks[rand.Intn(len(tracks))]

	fmt.Println("Opening:", url)

	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default: // Linux/others
		err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		fmt.Println("Error opening browser:", err)
	}
}
