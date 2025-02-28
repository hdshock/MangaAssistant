package models

// Series represents a manga series with its metadata
type Series struct {
	Series    string `json:"series"`
	Title     string `json:"title"`
	Volume    string `json:"volume"`
	Publisher string `json:"publisher"`
	Cover     string `json:"cover"`
	Volumes   int    `json:"volumes"`
	Chapters  int    `json:"chapters"`
}

type Manga struct {
	Title  string `json:"title"`
	Series string `json:"series"`
	Volume int    `json:"volume"`
	Cover  string `json:"cover"`
	Path   string `json:"path"`
}

type MangaLibrary struct {
	manga []Manga
}

func NewMangaLibrary() *MangaLibrary {
	return &MangaLibrary{
		manga: []Manga{},
	}
}

func (ml *MangaLibrary) AddManga(manga Manga) {
	ml.manga = append(ml.manga, manga)
}

func (ml *MangaLibrary) RemoveManga(manga Manga) {
	for i, m := range ml.manga {
		if m.Title == manga.Title && m.Volume == manga.Volume {
			ml.manga = append(ml.manga[:i], ml.manga[i+1:]...)
			break
		}
	}
}

func (ml *MangaLibrary) GetAllManga() []Manga {
	return ml.manga
}

func (ml *MangaLibrary) CheckDuplicates() []Manga {
	seen := make(map[string]bool)
	duplicates := []Manga{}
	for _, m := range ml.manga {
		key := m.Title + "-" + string(m.Volume)
		if seen[key] {
			duplicates = append(duplicates, m)
		} else {
			seen[key] = true
		}
	}
	return duplicates
} 