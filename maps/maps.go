package main

import "fmt"

func main() {
	trackList := make(map[int]string)
	trackList[1] = "Witch"
	trackList[2] = "You and Me and the Mountain"

	lyrics := map[string]string{
		"Witch": "You are a witch",
		"You and Me and the Mountain": "You and Me and the Mountain",
	}

	fmt.Println(trackList)
	fmt.Println(lyrics)
}
