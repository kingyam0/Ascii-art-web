package customfunctioninsidepackage

import (
	art "asciiweb/pkg/asciiart" //'art' is an ALIAS, if it wasnt present should use 'asciiart'
)

// GenerateArt will take the text and use the asciiart.go to genetate a
// string with all the converted ascii art characters

// example text => hi 	banner=> standard
// returns the ascii characters as a string...
// (currently it is printing on the terminal)

func GenerateArt(text string, style string) string {
	folderName := "banners/"

	filePath := folderName + style + ".txt"

	// contents of the banner.txt
	bannerContent := art.ReadLines(filePath)

	// Store all characters from 32 to 126 into a map
	bannerMap := art.StoreCharacters(bannerContent)

	// art.Print(text, bannerMap)

	return art.Store(text, bannerMap)
}

// TO:DO make the func above return a string instead of printing...
