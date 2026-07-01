package main

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
