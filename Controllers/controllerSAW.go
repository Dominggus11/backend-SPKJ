package controllers

import (
	"fmt"
	"math"
	"net/http"
	models "spkj/Models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func BeforeNormalisasi(UjianSekolah float64, RerataRaport float64, IPA float64, IPS float64, Minat string) (float64, float64, float64, float64, float64) {
	var (
		ci_ujian_sekolah float64
		ci_rerata        float64
		ci_ipa           float64
		ci_ips           float64
		ci_minat         float64
	)

	ci_ujian_sekolah = ConversiNilai(UjianSekolah)
	ci_rerata = ConversiNilai(RerataRaport)
	ci_ipa = ConversiNilai(IPA)
	ci_ips = ConversiNilai(IPS)
	ci_minat = ConversiJurusan(Minat)

	return ci_ujian_sekolah, ci_rerata, ci_ipa, ci_ips, ci_minat
}

func Normalisasi(c *gin.Context, siswas []models.Students) {
	var maxUjianSekolah float64
	var maxRerata float64
	var maxIPA float64
	var maxIPS float64
	var maxMinat float64

	maxUjianSekolah = 0.0
	maxRerata = 0.0
	maxIPA = 0.0
	maxIPS = 0.0
	maxMinat = 0.0

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
		if maxIPS < float64(siswa.Ci_IPS) {
			maxIPS = float64(siswa.Ci_IPS)
		}
		if maxMinat < float64(siswa.Ci_Minat) {
			maxMinat = float64(siswa.Ci_Minat)
		}
	}

	for _, siswa := range siswas {
		r_ujian_sekolah := siswa.Ci_UjianSekolah / maxUjianSekolah
		r_rerata := siswa.Ci_RerataRaport / maxRerata
		r_ipa := siswa.Ci_IPA / maxIPA
		r_ips := siswa.Ci_IPS / maxIPS
		r_minat := math.Round(siswa.Ci_Minat / maxMinat)

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
				RUjianSekolah: r_ujian_sekolah,
				RRerataRaport: r_rerata,
				RIpa:          r_ipa,
				RIps:          r_ips,
				RMinat:        r_minat,
			}
			db.Model(&input).Updates(student)
			c.JSON(http.StatusOK, gin.H{
				"data": input,
			})
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
				// fmt.Println(siswa.Nama, temp, kriteria.BobotKriteria, siswa.RUjianSekolah)
				temp = temp + (kriteria.BobotKriteria * siswa.RUjianSekolah)
			} else if kriteria.NamaKriteria == "Rerata Raport" {
				// fmt.Println(siswa.Nama, temp, kriteria.BobotKriteria, siswa.RRerataRaport)
				temp = temp + (kriteria.BobotKriteria * siswa.RRerataRaport)
			} else if kriteria.NamaKriteria == "Nilai IPA" {
				// fmt.Println(siswa.Nama, temp, kriteria.BobotKriteria, siswa.RIpa)
				temp = temp + (kriteria.BobotKriteria * siswa.RIpa)
			} else if kriteria.NamaKriteria == "Nilai IPS" {
				// fmt.Println(siswa.Nama, temp, kriteria.BobotKriteria, siswa.RIps)
				temp = temp + (kriteria.BobotKriteria * siswa.RIps)
			} else if kriteria.NamaKriteria == "Minat" {
				// fmt.Println(siswa.Nama, temp, kriteria.BobotKriteria, siswa.RMinat)
				temp = temp + (kriteria.BobotKriteria * siswa.RMinat)
			}
		}
		fmt.Println(siswa.Nama, temp)
	}
}
