package controllers

import (
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
)

func NormalisasiSMART(c *gin.Context, siswas []models.Students) {

	for _, siswa := range siswas {
		r_ujian_sekolah := (siswa.Ci_UjianSekolah - 1) / 4
		r_rerata := (siswa.Ci_RerataRaport - 1) / 4
		r_ipa := (siswa.Ci_IPA - 1) / 4
		r_ips := (siswa.Ci_IPS - 1) / 4
		// r_minat := (siswa.Ci_Minat - 1) / 4

		// input data normalisasi ujian sekolah
		db := models.DB
		// Get model if exist
		var input models.Students

		if err := db.Where("nisn = ?", siswa.NISN).First(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": " Data Siswa Tidak Tersedia"})
			return
		}

		err := db.Where("nisn = ?", siswa.NISN).First(&input).Error
		if err == nil {
			student := models.Students{
				RUjianSekolah_SMART: r_ujian_sekolah,
				RRerataRaport_SMART: r_rerata,
				RIpa_SMART:          r_ipa,
				RIps_SMART:          r_ips,
				// RMinat_SMART:        r_minat,
			}
			db.Model(&input).Updates(student)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Masuk"})
			return
		}
	}
}

func ResultSMART(c *gin.Context, criterias []models.Criterias, siswas []models.Students) {

	for _, siswa := range siswas {
		var temp float64 = 0.0
		for _, kriteria := range criterias {
			if kriteria.NamaKriteria == "Ujian Sekolah" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RUjianSekolah_SMART)
				}
			} else if kriteria.NamaKriteria == "Rerata Raport" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RRerataRaport_SMART)
				}
			} else if kriteria.NamaKriteria == "Nilai IPA" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RIpa_SMART)
				}
			} else if kriteria.NamaKriteria == "Nilai IPS" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RIps_SMART)
				}
			}

		}

		if temp > 74 {
			siswa.Jurusan_SMART = "IPA"
		} else {
			siswa.Jurusan_SMART = "IPS"
		}
		if temp == 0 {
			temp = 1
		}
		// input data normalisasi ujian sekolah
		db := models.DB
		// Get model if exist
		var input models.Students

		if err := db.Where("nisn = ?", siswa.NISN).First(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": " Data Siswa Tidak Tersedia"})
			return
		}

		err := db.Where("nisn = ?", siswa.NISN).First(&input).Error
		if err == nil {
			student := models.Students{
				ResultVi_SMART: temp,
				Jurusan_SMART:  siswa.Jurusan_SMART,
			}
			db.Model(&input).Updates(student)
			// c.JSON(http.StatusOK, gin.H{
			// 	"data": input,
			// })
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Masuk"})
			return
		}
	}
}
