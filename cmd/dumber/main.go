package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {

	var section string

	var headerCounters [7]int

	mdFilePath, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Println(err)
	}

	tmpFilePath := filepath.Dir(mdFilePath) + "/." + filepath.Base(mdFilePath) + ".tmp"

	mdFileHandler, err := os.Open(mdFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer mdFileHandler.Close()

	mdTmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		log.Println(err)
	}
	defer mdTmpFile.Close()

	headerLine := regexp.MustCompile(`^(#{1,6})\s+([\d\.]*)\s*(.*)$`)

	scanner := bufio.NewScanner(mdFileHandler)
	for scanner.Scan() {
		line := scanner.Text()
		matches := headerLine.FindStringSubmatch(line)

		if len(matches) == 4 {
			header := matches[1]
			currentHeaderType := len(matches[1])
			title := matches[3]
			headerCounters[currentHeaderType]++

			for headerType := 1; headerType <= 6; headerType++ {
				AddSectionChunk(&section, headerCounters[headerType], currentHeaderType, headerType)
			}

			_, err := io.WriteString(mdTmpFile, header+" "+section+" "+title+"\n")
			if err != nil {
				panic(err)
			}

			section = ""

			for i := currentHeaderType + 1; i <= 6; i++ {
				headerCounters[i] = 0
			}
		} else {
			_, err := io.WriteString(mdTmpFile, line+"\n")
			if err != nil {
				panic(err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = os.Rename(tmpFilePath, mdFilePath)
	if err != nil {
		log.Fatal(err)
	}

}

func AddSectionChunk(s *string, hc int, cht int, ht int) {
	if hc > 0 && cht >= ht {
		*s += strconv.Itoa(hc) + "."
	}
}
