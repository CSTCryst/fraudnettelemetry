package internal

import (
	"unicode/utf8"

	"golang.org/x/exp/slices"
)

// JavaScript - String.prototype.slice: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/slice
func SliceString(str string, indexStart, indexEnd int) string {
	strLen := len(str)
	if strLen == 0 {
		return ""
	}
	if indexStart < 0 {
		indexStart = strLen + indexStart
	}
	if indexEnd < 0 {
		indexEnd = strLen + indexEnd
	}
	if indexStart < 0 {
		indexStart = 0
	}
	if indexEnd > strLen {
		indexEnd = strLen
	}
	if indexStart > indexEnd {
		return ""
	}
	return str[indexStart:indexEnd]
}

func ExtractStrMapStrKeys(m map[string]string) []string {
	lenM := len(m)
	if lenM <= 0 {
		return nil
	}
	output := make([]string, 0, lenM)
	for k := range m {
		output = append(output, k)
	}
	slices.Sort(output)
	return output
}

// Source: https://stackoverflow.com/a/68124773
func CountTotalDigits(i int) int {
	if i >= 1e18 {
		return 19
	}
	count := 1
	for j := 10; j <= i; count++ {
		j *= 10
	}
	return count
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uintptr | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128
}

type Any interface {
	~string | Number
}

func IndexNum[T Number](r []T, i int) T {
	if (len(r) - 1) < i {
		return 0
	}
	return r[i]
}

func SumUnicodeValue[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uintptr | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64](s string) T {
	var t, value rune = 0, 0
	for i := range s {
		value, _ = utf8.DecodeRuneInString(s[i:])
		t += value
	}
	return T(t)
}
