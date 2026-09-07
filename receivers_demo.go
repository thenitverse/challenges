package main

import "fmt"

type Album struct {
	Title    string
	Stickers int
}

func (a Album) Rename(newTitle string) {
	a.Title = newTitle
}

func (a *Album) AddStickers(count int) {
	if a == nil {
		return
	}
	a.Stickers += count
}

func renameWithValue(album Album, newTitle string) {
	album.Rename(newTitle)
}

func renameWithPointer(album *Album, newTitle string) {
	album.Rename(newTitle)
}

func addWithValue(album Album, count int) {
	album.AddStickers(count)
}

func addWithPointer(album *Album, count int) {
	album.AddStickers(count)
}

func safeAddNilAlbum(count int) {
	var p *Album
	p.AddStickers(count)
}

func main() {
	fmt.Println("Running Album Receiver Tests...")

	// Test 1: Value receiver does not persist mutation
	album1 := Album{Title: "Animals", Stickers: 3}
	renameWithValue(album1, "Ocean")
	fmt.Printf("Test 1 (renameWithValue): Expected 'Animals', Got '%s'\n", album1.Title)

	// Test 2: Pointer receiver persists mutation
	album2 := Album{Title: "Space", Stickers: 4}
	addWithPointer(&album2, 2)
	fmt.Printf("Test 2 (addWithPointer): Expected 6, Got %d\n", album2.Stickers)

	// Test 3: Value passed to helper copies the struct, so modifications don't affect original
	album3 := Album{Title: "Sports", Stickers: 5}
	addWithValue(album3, 3)
	fmt.Printf("Test 3 (addWithValue): Expected 5, Got %d\n", album3.Stickers)

	// Test 4: Pointer receiver with value method still does not mutate original
	album4 := Album{Title: "Dinosaurs", Stickers: 7}
	renameWithPointer(&album4, "Robots")
	fmt.Printf("Test 4 (renameWithPointer): Expected 'Dinosaurs', Got '%s'\n", album4.Title)

	// Test 5: Safe handling of nil pointer receiver
	fmt.Println("Test 5 (safeAddNilAlbum): Calling method on nil pointer...")
	safeAddNilAlbum(10)
	fmt.Println("Test 5 Passed: No panic on nil receiver!")
}
