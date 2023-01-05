package controllers

import models "spkj/Models"

// get data criterias from database
func GetDataKriteria() []models.Criterias {
	var criterias []models.Criterias
	models.DB.Find(&criterias)
	// var criterias []models.Criterias
	return criterias
}

// get data siswas from database
func GetDataSiswa() []models.Students {
	var students []models.Students
	models.DB.Find(&students)
	var siswas []models.Students
	siswas = append(siswas, students...)
	return siswas
}
