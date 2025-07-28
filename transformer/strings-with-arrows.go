package transformer

import "strings"

// based off of https://github.com/davidcallanan/py-myopl-code/blob/master/ep14/strings_with_arrows.py
func StringsWithArrows(text string, bounds Bounds, linePrefix string) string {
	result := ""

	start := bounds.Start
	end := bounds.End

	idxStart := max(0, rfind(text, '\n', 0, int(start.Idx)))
	idxEnd := find(text, '\n', idxStart+1, len(text))
	if idxEnd < 0 {
		idxEnd = len(text)
	}

	lineCount := end.Ln - start.Ln + 1
	for i := range lineCount {
		line := text[idxStart:idxEnd]
		var colStart int16
		if i == 0 {
			colStart = start.Col
		} else {
			colStart = 0
		}
		var colEnd int16
		if i == lineCount-1 {
			colEnd = end.Col
		} else {
			colEnd = int16(len(line))
		}

		result += line + "\n"
		result += strings.Repeat(" ", int(colStart)) + strings.Repeat("^", int(colEnd-colStart))

		idxStart = idxEnd
		idxEnd = find(text, '\n', idxStart+1, len(text))
		if idxEnd < 0 {
			idxEnd = len(text)
		}
	}

	result = strings.ReplaceAll(result, "\t", "")
	result = linePrefix + result
	result = strings.ReplaceAll(result, "\n", "\n"+linePrefix)

	return result
}

func rfind(text string, char byte, start, end int) int {
	for i := end - 1; i >= start; i-- {
		if text[i] == char {
			return i
		}
	}
	return -1
}

func find(text string, char byte, start, end int) int {
	for i := start; i < end; i++ {
		if text[i] == char {
			return i
		}
	}
	return -1
}
