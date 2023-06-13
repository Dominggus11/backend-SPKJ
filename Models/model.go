package models

import (
	"time"

	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	ID                  uint
	Nama                string  `json:"nama"`
	NISN                string  `json:"nisn"`
	UjianSekolah        float64 `json:"ujian_sekolah"`
	RerataRaport        float64 `json:"rerata_raport"`
	IPA                 float64 `json:"ipa"`
	Minat               string  `json:"minat"`
	Ci_UjianSekolah     float64 `json:"ci_ujian_sekolah"`
	Ci_RerataRaport     float64 `json:"ci_rerata_raport"`
	Ci_IPA              float64 `json:"ci_ipa"`
	RUjianSekolah_SAW   float64 `json:"r_ujian_sekolah_saw"`
	RRerataRaport_SAW   float64 `json:"r_rerata_raport_saw"`
	RIpa_SAW            float64 `json:"r_ipa_saw"`
	ResultVi_SAW        float64 `json:"resultVi_saw"`
	Jurusan_SAW         string  `json:"jurusan_saw"`
	RUjianSekolah_SMART float64 `json:"r_ujian_sekolah_smart"`
	RRerataRaport_SMART float64 `json:"r_rerata_raport_smart"`
	RIpa_SMART          float64 `json:"r_ipa_smart"`
	ResultVi_SMART      float64 `json:"resultVi_smart"`
	Jurusan_SMART       string  `json:"jurusan_smart"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Criterias struct {
	gorm.Model
	ID            uint    `json:"id"`
	NamaKriteria  string  `json:"nama"`
	BobotKriteria float64 `json:"bobot"`
	Is_active     int     `json:"is_active"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Users struct {
	gorm.Model
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}
