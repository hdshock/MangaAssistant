package main

import (
	"archive/zip"
	"encoding/xml"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"manga-assistant/handlers"
	"manga-assistant/models"
)

// CORS middleware
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// ComicInfo represents the metadata in comicinfo.xml
type ComicInfo struct {
	Title    string `xml:"Title"`
	Series   string `xml:"Series"`
	Volume   int    `xml:"Volume"`
	Number   string `xml:"Number"`
	FilePath string `xml:"-"` // Path to the CBZ file
}

// Function to extract metadata from a CBZ file
func extractMetadataFromCBZ(path string) (ComicInfo, error) {
	info := ComicInfo{FilePath: path}

	// Open the CBZ file
	r, err := zip.OpenReader(path)
	if err != nil {
		return info, err
	}
	defer r.Close()

	// Look for comicinfo.xml
	for _, f := range r.File {
		if f.Name == "ComicInfo.xml" {
			rc, err := f.Open()
			if err != nil {
				return info, err
			}
			defer rc.Close()

			// Decode the XML
			if err := xml.NewDecoder(rc).Decode(&info); err != nil {
				return info, err
			}
			break
		}
	}

	return info, nil
}

// Function to parse chapter/episode number from a filename
func parseChapterNumber(filename string) string {
	// Regex to find numbers in the filename
	re := regexp.MustCompile(`\d+(\.\d+)?`)
	matches := re.FindAllString(filename, -1)

	// Exclude numbers that appear in every filename (e.g., "+99 stick chapter 02.cbz")
	if len(matches) > 1 {
		return matches[len(matches)-1] // Use the last number
	} else if len(matches) == 1 {
		return matches[0]
	}
	return ""
}

// Function to walk directories and return a list of manga files with metadata
func walkDirectory(root string) ([]ComicInfo, error) {
	var mangaFiles []ComicInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".cbz") {
			// Extract metadata from the CBZ file
			metadata, err := extractMetadataFromCBZ(path)
			if err != nil {
				log.Printf("Error extracting metadata from %s: %s\n", path, err)
				return nil
			}

			// If metadata is missing, fall back to folder and file names
			if metadata.Series == "" {
				metadata.Series = filepath.Base(filepath.Dir(path))
			}
			if metadata.Number == "" {
				metadata.Number = parseChapterNumber(info.Name())
			}

			mangaFiles = append(mangaFiles, metadata)
		}
		return nil
	})
	return mangaFiles, err
}

func main() {
	// Initialize the manga library
	library := models.NewMangaLibrary()

	// Define routes with CORS enabled
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Manga Assistant is running!")
	})
	http.HandleFunc("/api/manga", enableCORS(handlers.GetMangaList(library)))
	http.HandleFunc("/api/manga/add", enableCORS(handlers.AddManga(library)))
	http.HandleFunc("/api/manga/remove", enableCORS(handlers.RemoveManga(library)))
	http.HandleFunc("/api/manga/check-duplicates", enableCORS(handlers.CheckDuplicates(library)))
	http.HandleFunc("/api/manga/scrape", enableCORS(handlers.ScrapeMangaInfo))
	http.HandleFunc("/api/collections", enableCORS(handlers.GetCollections(library)))
	http.HandleFunc("/api/settings", enableCORS(handlers.GetSettings))
	http.HandleFunc("/api/libraries", enableCORS(handlers.GetLibraries))
	http.HandleFunc("/api/libraries/add", enableCORS(handlers.AddLibrary))
	http.HandleFunc("/api/libraries/remove", enableCORS(handlers.RemoveLibrary))
	http.HandleFunc("/scan", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		// Walk the directory and scan all subfolders
		mangaFiles, err := walkDirectory("/mnt/manga")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error scanning directory: %s", err)))
			return
		}
		fmt.Println("Scanned manga files:", mangaFiles)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Scan initiated successfully!"))
	}))
	http.HandleFunc("/browse", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		// List files in the base directory
		files, err := os.ReadDir("/mnt/manga")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error browsing directory: %s", err)))
			return
		}
		var fileNames []string
		for _, file := range files {
			fileNames = append(fileNames, file.Name())
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Files: %v", fileNames)))
	}))

	// Register the /api/series route only once
	log.Println("Registering /api/series route...")
	http.HandleFunc("/api/series", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling /api/series request...")
		// Scan the manga library
		libraryPath := os.Getenv("MANGA_LIBRARY_PATH")
		series, err := handlers.ScanMangaLibrary(libraryPath)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to scan manga library: %v", err), http.StatusInternalServerError)
			return
		}

		// Return the series as JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(series); err != nil {
			http.Error(w, fmt.Sprintf("Failed to encode series: %v", err), http.StatusInternalServerError)
		}
	}))

	// Add this route
	http.HandleFunc("/api/series/detail", enableCORS(handlers.GetSeriesDetail(library)))

	// Start the server
	log.Println("Starting Manga Assistant on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
} 