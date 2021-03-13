package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	var sectionString string

	var headerCounters [7]int

	var headerSection = regexp.MustCompile(`^(#{1,6})\s*(.*)$`)

	var tmpFileName = ".dumber_" + RandomString(12) + ".tmp"

	mdFile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer mdFile.Close()

	mdTmpFile, err := os.Create("/tmp/" + tmpFileName)
	if err != nil {
		log.Println(err)
	}
	defer mdTmpFile.Close()

	fmt.Println(tmpFileName)

	scanner := bufio.NewScanner(mdFile)
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
}

func AddSectionChunk(s *string, hc int, cht int, ht int) {
	if hc > 0 && cht >= ht {
		*s += strconv.Itoa(hc) + "."
	}
}

func RandomString(n int) string {
        rand.Seed(time.Now().UnixNano())
	var chars = []rune("abcde__fghijklmno_pqrstuvwxyzABC____DEFGHIJKLM___NOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}

	return string(s)
}
