package main

import (
	"fmt"
	"moviedl/scrapers/lk21"
)

func main() {
	// reGenre, err := regexp.Compile(`(^Nonton\s)([A-Za-z\s0-9]+)([\(\)0-9]+)`)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// finds := reGenre.FindAllString("Nonton Oasis Knebworth 1996 (2021) Film Subtitle Indonesia Streaming Movie Download", -1)
	// log.Println(finds)
	// fmt.Println(strings.Join(finds[:], ""))
	s := lk21.Download()
	// log.Println(finds)
	fmt.Println(s)
	fmt.Printf("\n\n")
}
