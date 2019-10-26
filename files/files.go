package files

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Metadata holds information about the file
type Metadata struct {
	gorm.Model
	Filename string
}

// Content has the contents of a file
type Content struct {
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
