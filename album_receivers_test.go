package main

import (
	"fmt"
	"testing"
)

func formatAlbum(album Album) string {
	return fmt.Sprintf("Title: %q, Stickers: %d", album.Title, album.Stickers)
}

func Test(t *testing.T) {
	type albumCase struct {
		name     string
		before   Album
		action   func(*Album)
		expected Album
	}

	runCases := []albumCase{
		{
			name:   "rename with value does not persist",
			before: Album{Title: "Animals", Stickers: 3},
			action: func(album *Album) {
				renameWithValue(*album, "Ocean")
			},
			expected: Album{Title: "Animals", Stickers: 3},
		},
		{
			name:   "add with pointer persists",
			before: Album{Title: "Space", Stickers: 4},
			action: func(album *Album) {
				addWithPointer(album, 2)
			},
			expected: Album{Title: "Space", Stickers: 6},
		},
		{
			name:   "add with value changes only the copy",
			before: Album{Title: "Sports", Stickers: 5},
			action: func(album *Album) {
				addWithValue(*album, 3)
			},
			expected: Album{Title: "Sports", Stickers: 5},
		},
	}

	submitCases := append(runCases, []albumCase{
		{
			name:   "rename with pointer still does not persist",
			before: Album{Title: "Dinosaurs", Stickers: 7},
			action: func(album *Album) {
				renameWithPointer(album, "Robots")
			},
			expected: Album{Title: "Dinosaurs", Stickers: 7},
		},
		{
			name:   "nil album add is safe",
			before: Album{Title: "Music", Stickers: 1},
			action: func(album *Album) {
				safeAddNilAlbum(10)
			},
			expected: Album{Title: "Music", Stickers: 1},
		},
		{
			name:   "pointer add happy path again",
			before: Album{Title: "Travel", Stickers: 8},
			action: func(album *Album) {
				addWithPointer(album, 4)
			},
			expected: Album{Title: "Travel", Stickers: 12},
		},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		album := test.before
		test.action(&album)

		if album != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Case: %s
Input before: %s

Expected after: %s
Actual after:   %s
Fail
`, test.name, formatAlbum(test.before), formatAlbum(test.expected), formatAlbum(album))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Case: %s
Input before: %s

Expected after: %s
Actual after:   %s
Pass
`, test.name, formatAlbum(test.before), formatAlbum(test.expected), formatAlbum(album))
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

var withSubmit = true
