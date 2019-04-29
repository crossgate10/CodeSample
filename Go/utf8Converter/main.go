package main

import (
	"io"
	"os"
	"strconv"

	"golang.org/x/text/encoding/traditionalchinese"
)

func main() {
	fileFolder := "/home/ben/TestData/TestConvert/"
	var files = []string{"big5-mac.xml", "big5-unix.xml", "big5-win.xml", "utf8-mac.xml", "utf8-unix.xml", "utf8-win.xml"}
	for i, f := range files {
		f, err := os.Open(fileFolder + f)
		if err != nil {
			// handle file open error
		}
		out, err := os.Create(fileFolder + "output" + strconv.Itoa(i) + ".xml")
		if err != nil {
			// handler error
		}

		r := traditionalchinese.Big5.NewDecoder().Reader(f)

		io.Copy(out, r)

		out.Close()
		f.Close()
	}
}
