package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) []string {
	file, _ := os.Open(path)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// StoreCharacters store the characters into a map with the key rune and values being the lines...
// lines are the lines from the banner file
func StoreCharacters(lines []string) map[rune][]string {
	bannerMap := make(map[rune][]string)

	// what is our final goal?
	// return a map whicb contains the key as a rune
	// with its values being the lines of the "asciiart"

	// we have 32 to 126 characters in the banner
	count := rune(31)

	for _, line := range lines {
		// when we have nothing
		// what should we do?
		if line == "" {
			count++
		} else {
			// add to map
			bannerMap[count] = append(bannerMap[count], line)
		}
	}

	return bannerMap
}

func Print(str string, banner map[rune][]string) {
	words := strings.Split(str, "\n")

	for _, word := range words {
		for i := 0; i < 8; i++ {
			line := ""
			for _, r := range word {
				line = line + banner[r][i]
			}
			fmt.Println(line)
		}
	}
}

// TO:DO

// func to make a string instead of printing

// Here

func Store(theinput string, banner map[rune][]string) string {
	convertedstr := "" // we are only declaring/initializing

	// always consider the scope

	words := strings.Split(theinput, "\r\n")

	for _, word := range words {
		for i := 0; i < 8; i++ {
			line := ""
			for _, r := range word {
				line = line + banner[r][i]
			}
			// we talked about the scope
			// fmt.Println(line)
			convertedstr = convertedstr + line + "\n"

		}
	}

	return convertedstr
}
