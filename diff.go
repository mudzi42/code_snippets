package codesnippets

// diff to string filespackage testdata

import "github.com/pmezard/go-difflib/difflib"

// Diff returns the diff between to strings
func Diff(a, b string) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(a),
		B:        difflib.SplitLines(b),
		FromFile: "Got",
		ToFile:   "Want",
		Context:  3,
	}
	text, _ := difflib.GetUnifiedDiffString(diff)
	return text
}
