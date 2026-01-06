package str

import (
	"iter"
	"slices"
	"strings"
)

func SplitLinesAndNormalizeSeq(s string, preserveEmpyLines bool) iter.Seq[string] {
	return func(yield func(string) bool) {
		for line := range strings.Lines(s) {
			trimmed := strings.TrimSpace(line)
			if !preserveEmpyLines && trimmed == "" {
				continue
			}
			if !yield(trimmed) {
				return
			}
		}
	}
}

func SplitLinesAndNormalize(s string, preserveEmptyLines bool) []string {
	return slices.Collect(SplitLinesAndNormalizeSeq(s, preserveEmptyLines))
}
