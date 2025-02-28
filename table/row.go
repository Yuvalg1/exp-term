package table

import (
	"fmt"
	"log"
	"mime"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetEntryContent(entry os.DirEntry, currentwd string) table.Row {
	return table.Row{
		entry.Name(),
		getFileType(entry),
		getSize(entry, currentwd),
		getDateModified(entry, currentwd),
	}
}

func formatSize(size int64) string {
	const MAX_SIZE_IN_BYTES int64 = 100000
	if size >= MAX_SIZE_IN_BYTES*1000*1000 {
		return fmt.Sprintf("%f gB", float64(size)/1000/1000/1000)
	}
	if size >= MAX_SIZE_IN_BYTES*1000 {
		return fmt.Sprintf("%f mB", float64(size)/1000/1000)
	}
	if size >= MAX_SIZE_IN_BYTES {
		return fmt.Sprintf("%f kB", float64(size)/1000)
	}
	return fmt.Sprintf("%d bytes", size)
}

func getSize(entry os.DirEntry, currentwd string) string {
	path := currentwd + "/" + entry.Name()
	if entry.IsDir() {
		dir, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		return fmt.Sprintf("%d items", len(dir))
	} else {
		file, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}

		return formatSize(file.Size())
	}
}

func getFileType(entry os.DirEntry) string {
	fileFormat := strings.Split(entry.Name(), ".")
	if entry.IsDir() {
		return "Folder"
	} else {
		fileType := mime.TypeByExtension("." + fileFormat[len(fileFormat)-1])
		fileType = strings.Split(fileType, "/")[0]
		caser := cases.Title(language.AmericanEnglish)
		fileType = caser.String(fileType)

		return fileType
	}
}

func getDateModified(entry os.DirEntry, currentwd string) string {
	path := currentwd + "/" + entry.Name()
	fileOrDir, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	return fileOrDir.ModTime().String()
}
