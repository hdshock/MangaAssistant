package handlers

import (
	"encoding/json"
	"net/http"
	"manga-assistant/models"
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetMangaList(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mangaList := library.GetAllManga()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mangaList)
	}
}

func AddManga(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manga models.Manga
		if err := json.NewDecoder(r.Body).Decode(&manga); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		library.AddManga(manga)
		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveManga(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var manga models.Manga
		if err := json.NewDecoder(r.Body).Decode(&manga); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		library.RemoveManga(manga)
		w.WriteHeader(http.StatusOK)
	}
}

func CheckDuplicates(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		duplicates := library.CheckDuplicates()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(duplicates)
	}
}

func ScrapeMangaInfo(w http.ResponseWriter, r *http.Request) {
	// Placeholder for scraping logic
	w.WriteHeader(http.StatusNotImplemented)
}

func GetSeries(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for series logic
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetCollections(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Placeholder for collections logic
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func GetSettings(w http.ResponseWriter, r *http.Request) {
	// Placeholder for settings logic
	w.WriteHeader(http.StatusNotImplemented)
}

type Library struct {
	Path string `json:"path"`
}

var libraries []Library

func GetLibraries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(libraries)
}

func AddLibrary(w http.ResponseWriter, r *http.Request) {
	var library Library
	if err := json.NewDecoder(r.Body).Decode(&library); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	libraries = append(libraries, library)
	w.WriteHeader(http.StatusCreated)
}

func RemoveLibrary(w http.ResponseWriter, r *http.Request) {
	var library Library
	if err := json.NewDecoder(r.Body).Decode(&library); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for i, lib := range libraries {
		if lib.Path == library.Path {
			libraries = append(libraries[:i], libraries[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}

// ComicInfo represents the metadata in a CBZ file
type ComicInfo struct {
	Series    string `xml:"Series"`
	Title     string `xml:"Title"`
	Volume    string `xml:"Volume"`
	Publisher string `xml:"Publisher"`
	Cover     string `xml:"Cover"`
}

// ScanMangaLibrary scans the manga library for CBZ files and returns a list of series
func ScanMangaLibrary(libraryPath string) ([]ComicInfo, error) {
	var series []ComicInfo

	err := filepath.Walk(libraryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a CBZ
		if strings.HasSuffix(strings.ToLower(path), ".cbz") {
			// Open the CBZ file
			reader, err := zip.OpenReader(path)
			if err != nil {
				return fmt.Errorf("failed to open CBZ file: %v", err)
			}
			defer reader.Close()

			// Look for ComicInfo.xml
			for _, file := range reader.File {
				if file.Name == "ComicInfo.xml" {
					// Read the ComicInfo.xml file
					rc, err := file.Open()
					if err != nil {
						return fmt.Errorf("failed to open ComicInfo.xml: %v", err)
					}
					defer rc.Close()

					// Parse the XML
					var comicInfo ComicInfo
					data, err := ioutil.ReadAll(rc)
					if err != nil {
						return fmt.Errorf("failed to read ComicInfo.xml: %v", err)
					}
					if err := xml.Unmarshal(data, &comicInfo); err != nil {
						return fmt.Errorf("failed to parse ComicInfo.xml: %v", err)
					}

					// Add the series to the list
					series = append(series, comicInfo)
					break
				}
			}
		}
		return nil
	})

	return series, err
}

// GetSeriesDetail returns detailed information about a specific series
func GetSeriesDetail(library *models.MangaLibrary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the series name from the URL query
		seriesName := r.URL.Query().Get("name")
		if seriesName == "" {
			http.Error(w, "Series name is required", http.StatusBadRequest)
			return
		}

		// Scan the manga library for the specific series
		libraryPath := os.Getenv("MANGA_LIBRARY_PATH")
		series, err := ScanMangaLibrary(libraryPath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan manga library: %v", err), http.StatusInternalServerError)
			return
		}

		// Find the series by name
		var seriesDetail models.Series
		for _, s := range series {
			if s.Series == seriesName {
				seriesDetail = models.Series{
					Series:    s.Series,
					Title:     s.Title,
					Volume:    s.Volume,
					Publisher: s.Publisher,
					Cover:     s.Cover,
					Volumes:   0, // Placeholder, update this with actual data
					Chapters:  0, // Placeholder, update this with actual data
				}
				break
			}
		}

		// Return the series detail as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(seriesDetail); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode series detail: %v", err), http.StatusInternalServerError)
		}
	}
} 