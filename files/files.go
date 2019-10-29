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
	filename string `json:"filename"`
	size     int64  `json:"size"`
}

// Content has the contents of a file
type Content struct {
	gorm.Model
	cpf                string `json:"cpf"`
	private            int    `json:"private"`
	incompleto         int    `json:"incompleto"`
	dataUltimaCompra   string `json:"data_ultima_compra"`
	ticketMedio        string `json:"ticket_medio"`
	ticketUltimaCompra string `json:"ticket_ultima_compra"`
	lojaMaisFrequente  string `json:"loka_mais_frequente"`
	lojaUltimaCompra   string `json:"loka_ultima_compra"`
	metadataID         uint   `json:"metadata_id"`
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

	// begin a transaction to reduce the time to save in the database
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

// GetContents returns all contents
func GetContents() (contents []Content) {

	db.DB.Find(&contents)

	return
}

// GetMetadatas Return all metadata
func GetMetadatas() (metas []Metadata) {

	db.DB.Find(&metas)

	return
}

// GetMetadata return one metadata
func GetMetadata(id int) (metadata Metadata) {

	db.DB.First(&metadata, id)

	return
}
