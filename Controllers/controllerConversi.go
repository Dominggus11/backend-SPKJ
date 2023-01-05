package controllers

func ConversiNilai(nilai float64) float64 {
	var ci float64
	if nilai > 91 {
		ci = 5
	} else if nilai > 80 {
		ci = 4
	} else if nilai > 70 {
		ci = 3
	} else if nilai > 60 {
		ci = 2
	} else {
		ci = 1
	}
	return ci
}

func ConversiJurusan(jurusan string) float64 {
	var ci float64
	if jurusan == "IPA" {
		ci = 5
	} else {
		ci = 1
	}
	return ci
}
