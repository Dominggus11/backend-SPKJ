package controllers

import (
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func BeforeNormalisasi(c *gin.Context, siswas []models.Students) {
	for _, siswa := range siswas {
		siswa.Ci_UjianSekolah = ConversiNilai(siswa.UjianSekolah)
		siswa.Ci_RerataRaport = ConversiNilai(siswa.RerataRaport)
		siswa.Ci_IPA = ConversiNilai(siswa.IPA)

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
				Ci_UjianSekolah: siswa.Ci_UjianSekolah,
				Ci_RerataRaport: siswa.Ci_RerataRaport,
				Ci_IPA:          siswa.Ci_IPA,
			}
			db.Model(&input).Updates(student)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Masuk"})
			return
		}
	}
}

func Normalisasi(c *gin.Context, siswas []models.Students) {
	var maxUjianSekolah float64
	var maxRerata float64
	var maxIPA float64

	maxUjianSekolah = 0.0
	maxRerata = 0.0
	maxIPA = 0.0

	// find max Criterias
	for _, siswa := range siswas {
		if maxUjianSekolah < float64(siswa.Ci_UjianSekolah) {
			maxUjianSekolah = float64(siswa.Ci_UjianSekolah)
		}
		if maxRerata < float64(siswa.Ci_RerataRaport) {
			maxRerata = float64(siswa.Ci_RerataRaport)
		}
		if maxIPA < float64(siswa.Ci_IPA) {
			maxIPA = float64(siswa.Ci_IPA)
		}
	}

	for _, siswa := range siswas {
		r_ujian_sekolah := siswa.Ci_UjianSekolah / maxUjianSekolah
		r_rerata := siswa.Ci_RerataRaport / maxRerata
		r_ipa := siswa.Ci_IPA / maxIPA
		db := models.DB
		var input models.Students

		if err := db.Where("nisn = ?", siswa.NISN).First(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": " Data Siswa Tidak Tersedia"})
			return
		}

		err := db.Where("nisn = ?", siswa.NISN).First(&input).Error
		if err == nil {
			student := models.Students{
				RUjianSekolah_SAW: r_ujian_sekolah,
				RRerataRaport_SAW: r_rerata,
				RIpa_SAW:          r_ipa,
			}
			db.Model(&input).Updates(student)

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data Tidak Masuk"})
			return
		}
	}
}

func ResultSAW(c *gin.Context, criterias []models.Criterias, siswas []models.Students) {

	for _, siswa := range siswas {
		var temp float64 = 0.0
		for _, kriteria := range criterias {
			if kriteria.NamaKriteria == "Ujian Sekolah" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RUjianSekolah_SAW)
				}

			} else if kriteria.NamaKriteria == "Rerata Raport" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RRerataRaport_SAW)
				}

			} else if kriteria.NamaKriteria == "Nilai IPA" {
				if kriteria.Is_active == 2 {
					temp = temp + 0
				} else {
					temp = temp + (kriteria.BobotKriteria * siswa.RIpa_SAW)
				}

			}

		}

		if temp > 74 {
			siswa.Jurusan_SAW = "IPA"
		} else {
			siswa.Jurusan_SAW = "IPS"
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
		// fmt.Println("Nilai Temp : ", temp)
		err := db.Where("nisn = ?", siswa.NISN).First(&input).Error
		if err == nil {
			student := models.Students{
				ResultVi_SAW: temp,
				Jurusan_SAW:  siswa.Jurusan_SAW,
			}
			// fmt.Println("Nilai Temp : ", temp)
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
