package models

import (
	"time"

	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	ID          uint
	Nama        string      `json:"nama" binding:"required"`
	NISN        string      `json:"nisn" binding:"required"`
	NilaiSiswas NilaiSiswas `gorm:"foreignKey:StudentID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NilaiSiswas struct {
	gorm.Model
	StudentID    uint
	UjianSekolah float32 `json:"ujian_sekolah"`
	RerataRaport float32 `json:"rerata_raport"`
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

// UjianSekolah    float32 `json:"ujian_sekolah" binding:"required"`
// RerataRaport    float32 `json:"rerata_raport" binding:"required"`
// IPA             float32 `json:"ipa" binding:"required"`
// IPS             float32 `json:"ips" binding:"required"`
// Minat           string  `json:"minat" binding:"required"`
// Ci_UjianSekolah float32 `json:"ci_ujian_sekolah"`
// Ci_RerataRaport float32 `json:"ci_rerata-raport"`
// Ci_IPA          float32 `json:"ci_ipa"`
// Ci_IPS          float32 `json:"ci_ips"`
// Ci_Minat        float32 `json:"ci_minat"`
// Jurusan         string  `json:"jurusan"`

// IPA          float32 `json:"ipa" binding:"required"`
// IPS          float32 `json:"ips" binding:"required"`
// Minat        float32 `json:"minat" binding:"required"`
// Jurusan      string  `json:"jurusan"`
