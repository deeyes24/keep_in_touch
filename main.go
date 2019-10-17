package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Sync Demo")
	source, err := ioutil.ReadFile("testdata/input.txt")
	if err != nil {
		fmt.Errorf("Unable to Read Source from input.txt file \n")
		fmt.Errorf("Qutting\n")
		return
	}

	destination, err2 := ioutil.ReadFile("testdata/output.txt")
	if err2 != nil {
		fmt.Errorf("Unable to Read Destination from output.txt file \n")
		fmt.Errorf("Qutting\n")
		return
	}

	changesMade := Diff(destination, source)

	Update(&destination, changesMade)

	err = ioutil.WriteFile("testdta/output.txt", destination, 0644)
	if err != nil {
		fmt.Errorf("Unable to update destination ")
		return
	}

}

func Diff(orig, changed []byte) map[uint64]byte {

	diff := make(map[uint64]byte)
	if len(changed) == 0 {
		diff[0] = ' '
	} else {

		for index := uint64(0); index < uint64(len(changed)); index++ {
			if index < uint64(len(orig)) {
				if orig[index] != changed[index] {
					diff[index] = changed[index]
				}
			} else {
				diff[index] = changed[index]
			}
		}
	}

	return diff
}

func Sync(src, dest *[]byte) {
	for i, s := range *src {
		(*dest)[i] = s
	}
}

func Update(orig *[]byte, diff map[uint64]byte) {

	if len(diff) == 0 {
		//Nothing to be done
		return
	}
	if len(diff) == 1 {
		if diff[0] == ' ' {
			*orig = make([]byte, 0)
			return
		}
	}

	for k, v := range diff {
		if k < uint64(len(*orig)) {
			(*orig)[k] = v
		} else {
			*orig = append(*orig, v)
		}
	}
}
