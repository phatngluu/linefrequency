package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	// read file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// new hash map
	m := make(map[string]int)
	scanner := bufio.NewScanner(file)

	// scan for each line
	for scanner.Scan() {
		line := scanner.Text()
		// check key exist
		if _, ok := m[line]; !ok {
			m[line] = 1
		} else {
			m[line] = m[line] + 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	} else {
		// convert map to slice
		type kv struct {
			Key   string
			Value int
		}
		var ss []kv
		for k, v := range m {
			ss = append(ss, kv{k, v})
		}

		// sort the slice
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value > ss[j].Value
		})

		// print result
		for _, kv := range ss {
			// 	// %4d	Pad with spaces (width 4, right justified)
			// %-4d	Pad with spaces (width 4, left justified)
			fmt.Printf("%-4d \t %s \n", kv.Value, kv.Key)
		}
	}
}

// Reference
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
// https://stackoverflow.com/questions/18695346/how-to-sort-a-mapstringint-by-its-values
// https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
