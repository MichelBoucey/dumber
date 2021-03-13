package main

import (
	"bufio"
	"io"
	"log"
	"os"
        "path"
        "path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var sectionString string

	var headerCounters [7]int

        mdFile,err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Println(err)
	}

        mdFileDir := path.Dir(mdFile)

        tmpFile := mdFileDir + "/." + os.Args[1] + ".tmp"

	mdFileHandler, err := os.Open(mdFile)
	if err != nil {
		log.Fatal(err)
	}
	defer mdFileHandler.Close()

	mdTmpFile, err := os.Create(tmpFile)
	if err != nil {
		log.Println(err)
	}
	defer mdTmpFile.Close()

	headerSection := regexp.MustCompile(`^(#{1,6})\s*(.*)$`)

	scanner := bufio.NewScanner(mdFileHandler)
	for scanner.Scan() {
		matches := headerSection.FindStringSubmatch(scanner.Text())
		currentHeaderType := len(matches[1])
		headerCounters[currentHeaderType]++

		for headerType := 1; headerType <= 6; headerType++ {
			AddSectionChunk(&sectionString, headerCounters[headerType], currentHeaderType, headerType)
		}

		_, err := io.WriteString(mdTmpFile, matches[1]+" "+strings.TrimRight(sectionString, ".")+" "+matches[2]+"\n")
		if err != nil {
			panic(err)
		}

		sectionString = ""

		for i := currentHeaderType + 1; i <= 6; i++ {
			headerCounters[i] = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = os.Rename(tmpFile, mdFile)
        if err != nil {
            log.Fatal(err)
        }

}

func AddSectionChunk(s *string, hc int, cht int, ht int) {
	if hc > 0 && cht >= ht {
		*s += strconv.Itoa(hc) + "."
	}
}

