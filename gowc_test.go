package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCountBytes(t *testing.T) {
	inputStr := "This has exactly 26 bytes!"
	input := bufio.NewScanner(strings.NewReader(inputStr))
	const expectedCount int64 = 26

	actualCount := countIn(input, Bytes)

	if actualCount != expectedCount {
		t.Fatalf(`countIn("%s", Bytes) = %d, instead of expected value of %d`,
			inputStr, actualCount, expectedCount)
	}
}

func TestCountWords(t *testing.T) {
	inputStr := "This has exactly 5 words!"
	input := bufio.NewScanner(strings.NewReader(inputStr))
	const expectedCount int64 = 5

	actualCount := countIn(input, Words)

	if actualCount != expectedCount {
		t.Fatalf(`countIn("%s", Words) = %d, instead of expected value of %d`,
			inputStr, actualCount, expectedCount)
	}
}

func TestCountLines(t *testing.T) {
	inputStr := "This\nhas\nexactly\n6\nlines!\n\n"
	input := bufio.NewScanner(strings.NewReader(inputStr))
	const expectedCount int64 = 6

	actualCount := countIn(input, Lines)

	if actualCount != expectedCount {
		t.Fatalf(`countIn("%s", Lines) = %d, instead of expected value of %d`,
			inputStr, actualCount, expectedCount)
	}
}

func TestCountChars(t *testing.T) {
	inputStr := "哈囉世界! This has exactly 36 characters!"
	input := bufio.NewScanner(strings.NewReader(inputStr))
	const expectedCount int64 = 37

	actualCount := countIn(input, Chars)

	if actualCount != expectedCount {
		t.Fatalf(`countIn("%s", Chars) = %d, instead of expected value of %d`,
			inputStr, actualCount, expectedCount)
	}
}
