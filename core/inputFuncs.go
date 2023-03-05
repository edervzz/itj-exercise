package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() ([]string, []string) {
	addresses := []string{}
	names := []string{}

	isError := false
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please enter addresses filename")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSuffix(text, "\r")
		readFile, err := os.Open(text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			address := fileScanner.Text()
			e := strings.Split(address, ",")
			isError = false
			if len(e) != 3 {
				fmt.Println("This file does not look like address file.")
				readFile.Close()
				isError = true
				break
			}
			addresses = append(addresses, address)
		}
		if isError {
			continue
		}
		readFile.Close()
		break
	}

	for {
		fmt.Println("Please enter names filename")
		reader = bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSuffix(text, "\r")
		readFile, err := os.Open(text)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			name := fileScanner.Text()
			e := strings.Split(name, ",")
			isError = false
			if len(e) != 1 {
				fmt.Println("This file does not look like drivers file.")
				readFile.Close()
				isError = true
				break
			}
			names = append(names, name)
		}
		if isError {
			continue
		}
		readFile.Close()
		break
	}

	return addresses, names
}
