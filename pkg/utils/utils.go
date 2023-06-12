package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"moro.io/wordlebot/pkg/logger"
)

func WordsFile(lang string) string {
	return "configs/" + lang + ".txt"
}

func GetWord(fileName string, nthWord int) (string, error) {
	logger.Debug(fmt.Sprintf("Retrieve the %d-th word from file %s", nthWord, fileName))

	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		logger.Fatalf("Error opening file: %v", err)
		return "", err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read and process each line
	lc := 1
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		lc++
		if lc > nthWord {
			break
		}
	}

	// Check for any errors encountered during scanning
	if err := scanner.Err(); err != nil {
		logger.Fatalf("Error scanning file: %v", err)
	}
	return line, err
}

func LineCounter(fileName string) int {
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := fi.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			panic("Couldn't read from file " + fileName)
		}
	}
}
