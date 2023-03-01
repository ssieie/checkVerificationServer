package utils

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var chineseChars []rune

func init() {
	file, err := os.Open("textData.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char >= 0x4e00 && char <= 0x9fa5 {
				chineseChars = append(chineseChars, char)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to scan file: %v", err)
	}
}

func GetRandomText() []rune {
	rand.Seed(time.Now().UnixNano())

	var result []rune
	for i := 0; i < rand.Intn(3)+2; i++ {
		result = append(result, chineseChars[rand.Intn(len(chineseChars))])
	}

	return result
}
