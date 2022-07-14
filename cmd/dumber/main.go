// dumber, a command line tool for numbering Mardown sections
//
// Copyright (c) 2021-2022 Michel Boucey
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS
// OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
// NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

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
)

func main() {

	version := "1.1.1"

	var headerCounters [7]int
	var mdTmpFile *os.File
	var newLine string
	var pathSep string
	var rewrittenLine string
	var section string

	switch runtime.GOOS {
	case "windows":
		pathSep = "\\"
		newLine = "\r\n"
	default:
		pathSep = "/"
		newLine = "\n"
	}

	helpFlag := flag.Bool("h", false, "Show help")
	removeFlag := flag.Bool("r", false, "Remove section numbers from the .md file")
	versionFlag := flag.Bool("v", false, "Show version")
	writeFlag := flag.Bool("w", false, "Write section numbers to the .md file (default to stdout)")

	flag.Parse()

	if *versionFlag == true {
		fmt.Println("dumber v" + version + newLine + "Copyright © 2021-2022 Michel Boucey" + newLine + "Released under 3-Clause BSD License")
		os.Exit(-1)
	}

	if len(os.Args) == 1 || len(flag.Args()) == 0 && *helpFlag == false {
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
		log.Println(err)
	}

	mdFileHandler, err := os.Open(mdFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if *writeFlag {
		mdTmpFile, err = os.CreateTemp(filepath.Dir(mdFilePath)+pathSep, ".dumber-*.tmp")
		if err != nil {
			log.Fatal(err)
		}
	}

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

			if *removeFlag {
				rewrittenLine = header + " " + title
			} else {
				rewrittenLine = header + " " + section + " " + title
			}

			if *writeFlag {
				_, err := io.WriteString(mdTmpFile, rewrittenLine+newLine)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Println(rewrittenLine)
			}

			section = ""

			for i := currentHeaderType + 1; i <= 6; i++ {
				headerCounters[i] = 0
			}

		} else {

			if *writeFlag {
				_, err := io.WriteString(mdTmpFile, line+newLine)
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

	mdFileHandler.Close()

	if *writeFlag {
		mdTmpFile.Close()

		err = os.Rename(mdTmpFile.Name(), mdFilePath)
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
