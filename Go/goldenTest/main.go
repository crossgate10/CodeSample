package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ToJSON(w io.Writer) error {
	var result = json.NewEncoder(w).Encode(&struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}{
		"Foo",
		"Bar",
	})

	return result
}

func main() {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	if err := ToJSON(w); err != nil {
		fmt.Fprintln(os.Stderr, "error writing json: %s", err)
	}
	w.Flush()
	fmt.Println(string(b.Bytes()))
}
