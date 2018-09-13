package codesnippets

//Sorting 2d slice [][]string (slice of slices)

import (
	"fmt"
	"sort"
	"strings"
)

type tableData [][]string

func (t tableData) Len() int      { return len(t) }
func (t tableData) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t tableData) Less(i, j int) bool { return strings.Compare(t[i][0], t[j][0]) == -1 }

func sliceSort() {
	fmt.Println("Hello, playground")

	td := tableData{
		{"Btemplate", "PASS", "hit1_leaf1_blahblah", "10"},
		{"Dtemplate", "FAIL", "hit1_leaf2_blahblah", "N/A"},
		{"Ctemplate", "PASS", "hit1_leaf2_blahblah", "0"},
		{"Atemplate", "N/A", "N/A", "N/A"},
	}

	fmt.Printf("%v\n", td)
	sort.Sort(td)
	fmt.Printf("%v\n", td)
}
