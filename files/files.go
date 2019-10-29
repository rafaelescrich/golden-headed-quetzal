package files

import (
	"bufio"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rafaelescrich/golden-headed-quetzal/db"
)

// Metadata holds information about the file
type Metadata struct {
	gorm.Model
	filename string
	size     int64
}

// Content has the contents of a file
type Content struct {
	gorm.Model
	cpf                string
	private            int
	incompleto         int
	dataUltimaCompra   string
	ticketMedio        string
	ticketUltimaCompra string
	lojaMaisFrequente  string
	lojaUltimaCompra   string
	metadataID         uint
}

// ResponseFiles creates a json response with files
type ResponseFiles struct {
	Count int        `json:"count"`
	Files []Metadata `json:"files"`
}

// ResponseContent creates a json response with contents
type ResponseContent struct {
	Count    int       `json:"count"`
	Contents []Content `json:"contents"`
}

func scanLines(file multipart.File) (lines []string) {

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}

func saveMetadata(f string, s int64) (id uint) {
	meta := Metadata{filename: f, size: s}

	db.DB.Create(&meta)

	return meta.ID
}

func saveContent(file multipart.File, id uint) (err error) {

	lines := scanLines(file)
	if err != nil {
		return
	}

	// lines -1 to remove the header
	contents := make([]Content, len(lines)-1)

	for i, line := range lines {
		if i == 0 {
			continue
		}
		str := strings.Fields(line)

		priv, err := strconv.Atoi(str[1])
		if err != nil {
			priv = 0
		}

		inc, err := strconv.Atoi(str[2])
		if err != nil {
			inc = 0
		}

		// i-1 because we skip the header of the file
		contents[i-1] = Content{
			cpf:                str[0],
			private:            priv,
			incompleto:         inc,
			dataUltimaCompra:   str[3],
			ticketMedio:        str[4],
			ticketUltimaCompra: str[5],
			lojaMaisFrequente:  str[6],
			lojaUltimaCompra:   str[7],
			metadataID:         id,
		}
	}

	// begin a transaction
	tx := db.DB.Begin()

	// do some database operations in the transaction (use 'tx' from this point, not 'db')
	for _, content := range contents {
		tx.Create(&content)
	}

	// Or commit the transaction
	tx.Commit()

	return nil
}

// Save file uploaded to database
func Save(filename string, size int64, file multipart.File) (err error) {

	id := saveMetadata(filename, size)

	err = saveContent(file, id)

	if err != nil {
		return
	}

	return
}
