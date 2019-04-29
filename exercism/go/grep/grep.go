// Package grep is string operation tools.
package grep

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Search find out the matching part of the content by pattern user input.
// -n Print the line numbers of each matching line.
// -l Print only the names of files that contain at least one matching line.
// -i Match line using a case-insensitive comparison.
// -v Invert the program -- collect all lines that fail to match the pattern.
// -x Only match entire lines, instead of lines that contain a match.
func Search(pattern string, flags, files []string) []string {
	var result = []string{}
	fs := strings.Join(flags, "")

	if len(files) != 0 {
		for _, fn := range files {
			content, e := ioutil.ReadFile(fn)
			checkErr(e)

			hasHeader := len(files) != 1
			isInsensitive := strings.Contains(fs, "-i")
			contents := strings.Split(string(content), "\n")
			for lineNumber, s := range contents {
				element := ""
				if strings.Contains(fs, "-l") {
					if strings.Contains(s, pattern) {
						result = append(result, fn)
						break
					}
				} else {
					element = addFileName(hasHeader, fn, element)
					if strings.Contains(fs, "-n") {
						element = checkEmpty(element) + strconv.Itoa(lineNumber+1)
					}
					if strings.Contains(fs, "-v") {
						if s != "" && !strings.Contains(s, pattern) {
							result = append(result, checkEmpty(element)+s)
						}
					} else {
						if strings.Contains(fs, "-x") {
							if isInsensitive {
								if strings.ToLower(s) == strings.ToLower(pattern) {
									result = append(result, checkEmpty(element)+s)
								}
							} else {
								if s == pattern {
									result = append(result, checkEmpty(element)+s)
								}
							}
						} else {
							if isInsensitive {
								if strings.Contains(strings.ToLower(s), strings.ToLower(pattern)) {
									result = append(result, checkEmpty(element)+s)
								}
							} else {
								if strings.Contains(s, pattern) {
									result = append(result, checkEmpty(element)+s)
								}
							}
						}
					}
				}
			}
		}
	}
	return result
}

func checkEmpty(element string) string {
	if element != "" && element[len(element)-1] != ':' {
		return element + ":"
	}
	return element
}

func addFileName(hasHeader bool, header, s string) string {
	if hasHeader {
		return header + ":" + s
	}
	return s
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
