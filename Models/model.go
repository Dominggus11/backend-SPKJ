package models

import (
	"time"

	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	ID              uint
	Nama            string  `json:"nama"`
	NISN            string  `json:"nisn"`
	UjianSekolah    float64 `json:"ujian_sekolah"`
	RerataRaport    float64 `json:"rerata_raport"`
	IPA             float64 `json:"ipa"`
	IPS             float64 `json:"ips"`
	Minat           string  `json:"minat"`
	Ci_UjianSekolah float64 `json:"ci_ujian_sekolah"`
	Ci_RerataRaport float64 `json:"ci_rerata_raport"`
	Ci_IPA          float64 `json:"ci_ipa"`
	Ci_IPS          float64 `json:"ci_ips"`
	Ci_Minat        float64 `json:"ci_minat"`
	RUjianSekolah   float64 `json:"r_ujian_sekolah"`
	RRerataRaport   float64 `json:"r_rerata_raport"`
	RIpa            float64 `json:"r_ipa"`
	RIps            float64 `json:"r_ips"`
	RMinat          float64 `json:"r_minat"`
	ResultVi        float64 `json:"resultVi"`
	Jurusan         string  `json:"jurusan"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Criterias struct {
	ID            uint
	NamaKriteria  string  `json:"nama_kriteria" binding:"required"`
	BobotKriteria float64 `json:"bobot_kriteria" binding:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Users struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}

// type DataSiswa struct {
// 	Message []struct {
// 		DeletedAt      interface{} `json:"DeletedAt"`
// 		ID             int         `json:"ID"`
// 		Nama           string      `json:"nama"`
// 		Nisn           string      `json:"nisn"`
// 		UjianSekolah   int         `json:"ujian_sekolah"`
// 		RerataRaport   int         `json:"rerata_raport"`
// 		Ipa            int         `json:"ipa"`
// 		Ips            int         `json:"ips"`
// 		Minat          string      `json:"minat"`
// 		CiUjianSekolah int         `json:"ci_ujian_sekolah"`
// 		CiRerataRaport int         `json:"ci_rerata-raport"`
// 		CiIpa          int         `json:"ci_ipa"`
// 		CiIps          int         `json:"ci_ips"`
// 		CiMinat        int         `json:"ci_minat"`
// 		RUjianSekolah  int         `json:"r_ujian_sekolah"`
// 		RRerataRaport  int         `json:"r_rerata-raport"`
// 		RIpa           int         `json:"r_ipa"`
// 		RIps           int         `json:"r_ips"`
// 		RMinat         int         `json:"r_minat"`
// 		Jurusan        string      `json:"jurusan"`
// 		CreatedAt      time.Time   `json:"CreatedAt"`
// 		UpdatedAt      time.Time   `json:"UpdatedAt"`
// 	} `json:"message"`
// 	Response string `json:"response"`
// }
