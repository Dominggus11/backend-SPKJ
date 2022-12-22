package models

import (
	"time"
)

type Students struct {
	ID   uint
	Nama string `json:"nama"`
	// Description string `json:"description"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Criterias struct {
	ID            uint
	NamaKriteria  string  `json:"nama_kriteria" binding:"required"`
	BobotKriteria float32 `json:"bobot_kriteria" binding:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Users struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}
