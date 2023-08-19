//
// dumber, a command line tool for numbering and denumbering Mardown sections file
//
//      Copyright (c) 2021-2023 Michel Boucey
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
//      1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
//      2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer
//         in the documentation and/or other materials provided with the distribution.
//
//      3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived
//         from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED
// TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PRO-
// CUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
// ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

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
	"runtime"
	"strconv"
	"strings"
)

func main() {

	version := "2.0.0"

	var headerCounters [7]int
	var mdTmpFile *os.File
	var newLine string
	var mdLines []string
	var pathSep string
	var rewrittenLine string
	var section string
	var tocLines []string

	switch runtime.GOOS {
	case "windows":
		pathSep = "\\"
		newLine = "\r\n"
	default:
		pathSep = "/"
		newLine = "\n"
	}

	helpFlag := flag.Bool("h", false, "Show help")
	removeFlag := flag.Bool("r", false, "Remove table of contents and section numbers from the .md file")
	tocFlag := flag.Bool("t", false, "Add a table of contents to the .md file (can not be combined with -r)")
	versionFlag := flag.Bool("v", false, "Show version")
	writeFlag := flag.Bool("w", false, "Write section numbers to the .md file (default to stdout)")

	flag.Parse()

	if *versionFlag == true {
		fmt.Println("dumber v" + version + newLine + "Copyright © 2021-2023 Michel Boucey" + newLine + "Released under 3-Clause BSD License")
		os.Exit(-1)
	}

	if len(os.Args) == 1 || len(flag.Args()) == 0 && *helpFlag == false || *removeFlag == true && *tocFlag == true {
		fmt.Println("See -h for help")
		os.Exit(-1)
	}

	if *helpFlag {
		fmt.Println("Usage: dumber [OPTION] FILE" + newLine)
		flag.PrintDefaults()
		fmt.Println("")
		os.Exit(-1)
	}

	mdFilePath, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	mdFileHandler, err := os.Open(mdFilePath)
	if err != nil {
		log.Fatal(err)
	}

	tocLine := regexp.MustCompile(`^\s*-\s\[[\d\.]*\]\(#[\d\.]*\)`)
	headerLine := regexp.MustCompile(`^(#{1,6})\s+\[?([\d\.]*)(?:\]\(#\)\{name=[\d\.]*\})?\s*(.*)$`)

	scanner := bufio.NewScanner(mdFileHandler)
	for scanner.Scan() {

		line := scanner.Text()
		matches := headerLine.FindStringSubmatch(line)

		if len(matches) == 4 {

			header := matches[1]
			currentHeaderType := len(matches[1])
			title := matches[3]
			headerCounters[currentHeaderType]++

			if *removeFlag {

				rewrittenLine = header + " " + title

			} else {

				for headerType := 1; headerType <= 6; headerType++ {

					addSectionChunk(&section, headerCounters[headerType], currentHeaderType, headerType)

				}

				rewrittenLine = header + " " + section + " " + title
			}

			if *tocFlag {

				tocLines = append(tocLines, rewrittenLine)

			}

			mdLines = append(mdLines, rewrittenLine)

			if !*removeFlag {

				for i := currentHeaderType + 1; i <= 6; i++ {
					headerCounters[i] = 0
				}

				section = ""

			}

		} else {

			if !tocLine.Match([]byte(line)) {

				mdLines = append(mdLines, line)

			}

		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	mdFileHandler.Close()

	if *writeFlag {

		mdTmpFile, err = os.CreateTemp(filepath.Dir(mdFilePath)+pathSep, ".dumber-*.tmp")
		if err != nil {
			log.Fatal(err)
		}

		if *tocFlag {

			for _, line := range tocLines {

				matches := headerLine.FindStringSubmatch(line)

				_, _ = io.WriteString(mdTmpFile, strings.Repeat("    ", len(matches[1])-1)+"- "+"["+matches[2]+"](#"+matches[2]+") "+matches[3]+newLine)

			}

		}

		for _, line := range mdLines {

			matches := headerLine.FindStringSubmatch(line)

			if len(matches) == 4 && *tocFlag {

				_, _ = io.WriteString(mdTmpFile, matches[1]+" ["+matches[2]+"](#){name="+matches[2]+"} "+matches[3]+newLine)

			} else {

				_, _ = io.WriteString(mdTmpFile, line+newLine)

			}
		}

		mdTmpFile.Close()

		err = os.Rename(mdTmpFile.Name(), mdFilePath)
		if err != nil {
			log.Fatal(err)
		}

	} else {

		if *tocFlag {

			for _, line := range tocLines {

				matches := headerLine.FindStringSubmatch(line)

				fmt.Println(strings.Repeat("    ", len(matches[1])-1) + "- " + "[" + matches[2] + "](#" + matches[2] + ") " + matches[3])

			}

		}

		for _, line := range mdLines {

			matches := headerLine.FindStringSubmatch(line)

			if len(matches) == 4 && *tocFlag {

				fmt.Println(matches[1] + " [" + matches[2] + "](#){name=" + matches[2] + "} " + matches[3])

			} else {

				fmt.Println(line)

			}

		}
	}
}

func addSectionChunk(s *string, hc int, cht int, ht int) {
	if hc > 0 && cht >= ht {
		*s += strconv.Itoa(hc) + "."
	}
}
