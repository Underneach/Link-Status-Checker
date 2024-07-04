package src

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputData() {
	for {
		fmt.Print("Enter file with links path: ")
		linksFile, _ := input.ReadString('\n')
		linksFile = strings.TrimSpace(linksFile)
		file, err := os.OpenFile(linksFile, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Printf("Error in file reading : %d\n", err)
			continue
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			text = strings.TrimSpace(text)
			if text != "" {
				linksList = append(linksList, text)
			}
		}
		_ = file.Close()
		totalLinks = len(linksList)
		fmt.Printf("Total links: %d\n", totalLinks)
		if totalLinks == 0 {
			continue
		} else {
			break
		}
	}

	for {
		fmt.Print("Enter threads count: ")
		thRaw, _ := input.ReadString('\n')
		thRaw = strings.TrimSpace(thRaw)
		if threads, err = strconv.Atoi(thRaw); thRaw == "" || err != nil || threads < 1 {
			continue
		}
		break
	}
}
