package main

import (
	"fmt"
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

func main(){
	display("head")
}


func display(output string){
	switch output {
		case "head":
			fmt.Printf(asciiHead)
		case "tails":
			fmt.Printf(asciiTails)
		default:
			fmt.Println("Soemthing Went Wrong")
	}
	
}