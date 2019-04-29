package main

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func TestToJSON(t *testing.T) {
	testtable := []struct {
		tname string
	}{
		{
			tname: "ok",
		},
	}
	for _, tc := range testtable {
		t.Run(tc.tname, func(t *testing.T) {
			var b bytes.Buffer
			w := bufio.NewWriter(&b)
			err := ToJSON(w)
			if err != nil {
				t.Fatalf("failed writing json: %s", err)
			}
			w.Flush()
			gp := filepath.Join("testdata", filepath.FromSlash(t.Name())+".golden")
			if *update {
				t.Log("update golden file")
				if err := ioutil.WriteFile(gp, b.Bytes(), 0644); err != nil {
					t.Fatalf("failed to update golden file: %s", err)
				}
			}
			g, err := ioutil.ReadFile(gp)
			if err != nil {
				t.Fatalf("failed reading .golden: %s", err)
			}
			t.Log(string(b.Bytes()))
			if !bytes.Equal(b.Bytes(), g) {
				t.Errorf("bytes do not match .golden file")
			}
		})
	}
}
