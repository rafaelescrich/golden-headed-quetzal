package files

import (
	"time"

	"github.com/jinzhu/gorm"
)

// File metadata
type File struct {
	gorm.Model
	Filename string
}

// FileContent has the contents of a file
type FileContent struct {
	gorm.Model
	CPF                string
	Private            int
	Incompleto         int
	DataUltimaCompra   time.Time
	TicketMedio        float32
	TicketUltimaCompra float32
	LojaMaisFrequente  string
	LojaUltimaCompra   string
}
