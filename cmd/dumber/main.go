package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {

	version := "1.0.0"

	var section string
	var headerCounters [7]int

	versionFlag := flag.Bool("v", false, "Show version")
	helpFlag := flag.Bool("h", false, "Show help")
	writeFlag := flag.Bool("w", false, "Rewrite the md file (default to stdout)")
	// removeFlag := flag.Bool("r", false, "Remove sections numbers in the md file")

	flag.Parse()

	if *versionFlag == true {
		fmt.Println("dumber v" + version + "\nCopyright © 2021 Michel Boucey\nReleased under 3-Clause BSD License")
		os.Exit(-1)
	}

	if len(os.Args) == 1 || len(flag.Args()) == 0 && *helpFlag == false {
		fmt.Println("See -h for help)")
		os.Exit(-1)
	}

	if *helpFlag {
		fmt.Println("Usage: dumber [-w|-r] file.md\n")
		flag.PrintDefaults()
		fmt.Println("")
		os.Exit(-1)
	}

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

			if *writeFlag {
				_, err := io.WriteString(mdTmpFile, header+" "+section+" "+title+"\n")
				if err != nil {
					panic(err)
				}
			} else {

				fmt.Println(header + " " + section + " " + title)

			}

			section = ""

			for i := currentHeaderType + 1; i <= 6; i++ {
				headerCounters[i] = 0
			}
		} else {
			if *writeFlag {
				_, err := io.WriteString(mdTmpFile, line+"\n")
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println(line)

			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *writeFlag {
		err = os.Rename(tmpFilePath, mdFilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func AddSectionChunk(s *string, hc int, cht int, ht int) {
	if hc > 0 && cht >= ht {
		*s += strconv.Itoa(hc) + "."
	}
}
