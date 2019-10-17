package main

import (
	"io/ioutil"
	"testing"
)

var inputFile = "testdata/input.txt"
var outputFile = "testdata/output.txt"

func TestReadFile(t *testing.T) {

	byteFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("Unable to read file ")
	}
	if len(byteFile) == 0 {
		t.Errorf("Empty File Found")
	}

}

func TestWriteToFile(t *testing.T) {
	err := ioutil.WriteFile(outputFile, []byte{}, 0644)
	if err != nil {
		t.Errorf("Unable to write to file")
	}
}

func TestSyncSingleChar(t *testing.T) {
	var src = []byte{'a'}
	var dest = []byte{'b'}
	Sync(&src, &dest)
	for index := 0; index < len(src); index++ {
		if src[index] != dest[index] {
			t.Errorf("Found mismatch at %d for src is %d and dest is %d", index, src[index], dest[index])
		}
	}

}
func TestCompareByteArray(t *testing.T) {
	var orig = []byte{'a'}
	var changed = []byte{'b'}
	//Should return that at pos 0, we should replace by b
	diff := Diff(orig, changed)
	if diff[0] != 'b' {
		t.Errorf("Should have been b")
	}
	orig = []byte{'a'}
	changed = []byte{'a'}
	diff = Diff(orig, changed)
	if len(diff) != 0 {
		t.Errorf("There should be no diff")
	}

	orig = []byte{'a'}
	changed = []byte{' ', 'b'}
	diff = Diff(orig, changed)
	if diff[1] != 'b' {
		t.Errorf("Should have been b")
	}

	if diff[0] != ' ' {
		t.Errorf("Should have been whitespace")
	}
	orig = []byte{'a'}
	changed = []byte{}
	diff = Diff(orig, changed)
	if diff[0] != ' ' {
		t.Errorf("Should have been white space")
	}

}

func TestUpdateDestinatonData(t *testing.T) {
	//a, ab, a
	var orig = []byte{'a'}
	var changed = []byte{'a', 'b'}
	diff := Diff(orig, changed)

	Update(&orig, diff)

	if orig[0] != 'a' {
		t.Errorf("After updating original string ,1nd byte should be a")
	}

	if orig[1] != 'b' {
		t.Errorf("After updating original string ,2nd byte should be b")
	}

	orig = []byte{'b'}
	changed = []byte{'a'}
	diff = Diff(orig, changed)
	Update(&orig, diff)
	if orig[0] != 'a' {
		t.Errorf("After updating original string ,1nd byte should be a")
	}

	orig = []byte{'b'}
	changed = []byte{}
	diff = Diff(orig, changed)
	Update(&orig, diff)
	if len(orig) != 0 {
		t.Errorf("Length should  have been zero as we are removing chars Orig %q Changed %q   Diff Map %q   Orig %s", orig, changed, diff, orig)
	}

	orig = []byte{'b', 'd', 'a'}
	changed = []byte{}
	diff = Diff(orig, changed)
	Update(&orig, diff)
	if len(orig) != 0 {
		t.Errorf("Length should  have been zero as we are removing chars Orig %q Changed %q   Diff Map %q   Orig %s", orig, changed, diff, orig)
	}

}
