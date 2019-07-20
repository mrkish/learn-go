package maps

import (
	"errors"
	"fmt"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

func main() {
	trackList := make(map[int]string)
	trackList[1] = "Witch"
	trackList[2] = "You and Me and the Mountain"

	lyrics := map[string]string{
		"Witch":                       "You are a witch",
		"You and Me and the Mountain": "You and Me and the Mountain",
	}

	fmt.Println(trackList)
	fmt.Println(lyrics)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, value string) error {
	d[key] = value
	return nil
}
