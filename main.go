package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/briandowns/spinner"
)

const asciiHead = `
_______  _______  ______
|\     /|(  ____ \(  ___  )(  __  \
| )   ( || (    \/| (   ) || (  \  )
| (___) || (__    | (___) || |   ) |
|  ___  ||  __)   |  ___  || |   | |
| (   ) || (      | (   ) || |   ) |
| )   ( || (____/\| )   ( || (__/  )
|/     \|(_______/|/     \|(______/

`

const asciiTails = `
_________ _______ _________ _       _______
\__   __/(  ___  )\__   __/( \     (  ____ \
   ) (   | (   ) |   ) (   | (     | (    \/
   | |   | (___) |   | |   | |     | (_____
   | |   |  ___  |   | |   | |     (_____  )
   | |   | (   ) |   | |   | |           ) |
   | |   | )   ( |___) (___| (____/Y\____) |
   )_(   |/     \|\_______/(_______|_______)

`

func main() {
	s := spinner.New(spinner.CharSets[17], 100*time.Millisecond)
	s.Start()
	time.Sleep(2 * time.Second)
	s.Stop()
	
	output := getRandomOutput()
	display(output)
}

func display(output string) {
	switch output {
	case "head":
		fmt.Printf(asciiHead)
	case "tails":
		fmt.Printf(asciiTails)
	default:
		fmt.Println("Soemthing Went Wrong")
	}

}

func getRandomOutput() string  {
	if rand.Intn(2) == 0 {
		return "head"
	} else {
		return "tails"
	}
}
